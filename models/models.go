package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:!@#$%^&*(0)@tcp(106.187.54.95:3306)/ablog?charset=utf8&parseTime=true&loc=Local")
	orm.RegisterModelWithPrefix("t_", new(User), new(Article), new(Tag), new(FkArticleTag))
	orm.Debug = true
}

const (
	Publish = iota
	Draft
	Recycle
)

//用户
type User struct {
	Id  int `orm:"pk"`
	Usn string
	Pwd string
}

//文章
type Article struct {
	Id       int
	Title    string
	Content  string
	Writer   string
	Top      bool
	State    int8
	UrlName  string
	Pv       int
	Tags     string `orm:"size(250)"`
	CreateOn time.Time
	UpdateOn time.Time
}

func (this *Article) TagNames() []string {
	this.Tags = strings.TrimSpace(this.Tags)
	if this.Tags == "" {
		return make([]string, 0)
	}

	return strings.Split(this.Tags, ",")
}

//标签
type Tag struct {
	Id    int `orm:"pk"`
	Name  string
	Count int
}

//文章标签关联
type FkArticleTag struct {
	Id        int
	ArticleId int
	TagId     int
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GetOneUserByUsn(usn string) (user *User, err error) {
	o := orm.NewOrm()

	user = new(User)
	user.Usn = usn

	if err = o.Read(user, "usn"); err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return
}

func AddArticle(article *Article) (err error) {
	o := orm.NewOrm()
	if err = o.Begin(); err != nil {
		return
	}

	article.CreateOn = time.Now()
	article.UpdateOn = article.CreateOn

	_, err = o.Insert(article)

	fKArticleTags := make([]FkArticleTag, 0)
	tagNames := article.TagNames()
	if len(tagNames) > 10 {
		o.Rollback()
		return errors.New("标签数最多10个")
	}
	for _, v := range tagNames {
		if len(v) > 25 {
			o.Rollback()
			return errors.New("标签字符数最多25个")
		}

		tag := &Tag{Name: v}
		var id int64
		if _, id, err = o.ReadOrCreate(tag, "Name"); err != nil {
			o.Rollback()
			return err
		}

		tag.Id = int(id)
		tag.Count++

		if _, err = o.Update(tag, "Count"); err != nil {
			o.Rollback()
			return err
		}
		fKArticleTags = append(fKArticleTags, FkArticleTag{ArticleId: article.Id, TagId: tag.Id})
	}

	if len(fKArticleTags) > 0 {
		if _, err = o.InsertMulti(len(fKArticleTags), fKArticleTags); err != nil {
			o.Rollback()
			return err
		}
	}

	err = o.Commit()
	return
}
func UpdateArticle(article *Article) (err error) {
	o := orm.NewOrm()

	_, err = o.Update(article)

	return
}
func UpdateArticlesState(ids *[]int, state int8) (err error) {
	orm := orm.NewOrm()

	_, err = orm.Raw(`
		UPDATE t_article set state = ? 
		WHERE id in (?)
	`, state, ids).Exec()

	return
}
func DeleteArticles(ids *[]int) (err error) {
	o := orm.NewOrm()

	if err = o.Begin(); err != nil {
		return
	}

	//文章标签 count - 1
	_, err = o.Raw(`
		UPDATE t_tag AS t 
			INNER JOIN t_fk_article_tag AS fk ON fk.tag_id = t.id
			INNER JOIN t_article AS a ON a.id = fk.article_id
		SET count = count-1
		WHERE a.id IN (?) 
	`, ids).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	//删除count数0的标签
	_, err = o.Raw(`
		DELETE FROM t_tag
		WHERE count = 0 
	`).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	//删除文章和标签关联
	_, err = o.Raw(`
		DELETE 
			t_article, 
			t_fk_article_tag 
		FROM t_article
			INNER JOIN t_fk_article_tag ON t_fk_article_tag.article_id = t_article.id
		WHERE t_article.id IN (?)
	`, ids).Exec()
	if err != nil {
		o.Rollback()
		return
	}

	if err = o.Commit(); err != nil {
		o.Rollback()
		return
	}

	return
}
func GetAllArticle(title string, writer string, tag string, state int8, psize int, pinx int) (articles []*Article, err error) {
	orm := orm.NewOrm()

	if title = strings.TrimSpace(title); title != "" {
		_, err = orm.QueryTable("t_article").Filter("state", state).Filter("title__icontains", title).Limit(psize, psize*(pinx-1)).All(&articles)
	} else if writer = strings.TrimSpace(writer); writer != "" {
		_, err = orm.QueryTable("t_article").Filter("state", state).Filter("writer__icontains", writer).Limit(psize, psize*(pinx-1)).All(&articles)
	} else if tag = strings.TrimSpace(tag); tag != "" {
		_, err = orm.Raw(`
			SELECT
			a.*
			FROM
			t_article a
			INNER JOIN t_pk_article_tag m ON a.id = m.article_id
			INNER JOIN t_tag t ON m.tag_id = t.id
			WHERE t.name LIKE ? AND a.state = ?
			GROUP BY a.id
			ORDER BY a.create_on DESC
			LIMIT ? OFFSET ?
		`, fmt.Sprintf("%%%s%%", tag), state, psize, psize*(pinx-1)).QueryRows(&articles)
	} else {
		_, err = orm.QueryTable("t_article").Filter("state", state).Limit(psize, psize*(pinx-1)).All(&articles)
	}

	return
}
func GetHomeArticle(psize, pinx int) (articles []Article, err error) {
	o := orm.NewOrm()

	_, err = o.QueryTable("t_article").Filter("State", 0 /*已发布的*/).OrderBy("-CreateOn").Limit(psize, psize*(pinx-1)).All(&articles)

	return
}
func GetOneArticle(id int) (article *Article, err error) {
	o := orm.NewOrm()

	article = new(Article)
	article.Id = id

	if err = o.Read(article); err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return
}
func GetArticlesByTag(tag string) (article []Article, err error) {
	o := orm.NewOrm()

	_, err = o.Raw(`
		SELECT t_article.*
		FROM t_article
		INNER JOIN t_fk_article_tag ON t_fk_article_tag.article_id = t_article.id
		INNER JOIN t_tag ON t_tag.id = t_fk_article_tag.tag_id
		WHERE t_article.state = 0 AND t_tag.name = ?
		ORDER BY t_article.update_on DESC
	`, tag).QueryRows(&article)

	return
}
func GetCountByArticle(state int8) (count int64, err error) {
	o := orm.NewOrm()

	count, err = o.QueryTable("t_article").Filter("state", state).Count()

	return
}

//文章PV加一
func IncrArticlePv(id int) (err error) {
	o := orm.NewOrm()

	_, err = o.QueryTable("t_article").Filter("id", id).Update(orm.Params{"pv": orm.ColValue(orm.Col_Add, 1)})

	return
}

//获取文章归档
func GetArticleArchives(tag string) (rst map[string][]*Article, years []string, err error) {
	o := orm.NewOrm()

	rst = make(map[string][]*Article)
	years = make([]string, 0)
	articles := make([]*Article, 0)

	tag = strings.TrimSpace(tag)

	if tag == "" {
		_, err = o.QueryTable("t_article").Filter("state", 0).OrderBy("-create_on").All(&articles, "id", "title", "tags", "create_on")
	} else {
		_, err = o.Raw(`
			SELECT 
				a.id,
				a.title,
				a.tags,
				a.create_on
			FROM t_article AS a 
				INNER JOIN t_fk_article_tag AS fk ON fk.article_id = a.id
				INNER JOIN t_tag AS t ON t.id = fk.tag_id
			WHERE 
				a.state = 0
				AND t.name = ?
			ORDER BY create_on DESC
		`, tag).QueryRows(&articles)
	}

	if err != nil {
		return
	}

	for _, v := range articles {
		year := v.CreateOn.Format("2006")
		if _, ok := rst[year]; !ok {
			rst[year] = make([]*Article, 0)
			years = append(years, year)
		}

		rst[year] = append(rst[year], v)
	}

	return
}

func GetAllTag() (tags []Tag, err error) {
	orm := orm.NewOrm()

	_, err = orm.QueryTable("t_tag").All(&tags)

	return
}
