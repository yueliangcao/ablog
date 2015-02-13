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
	return strings.Split(strings.TrimRight(this.Tags, ","), ",")
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
		if _, _, err = o.ReadOrCreate(tag, "Name"); err != nil {
			o.Rollback()
			return err
		}

		tag.Count++

		o.Update(tag, "Count")
		fKArticleTags = append(fKArticleTags, FkArticleTag{ArticleId: article.Id, TagId: tag.Id})
	}
	if _, err = o.InsertMulti(len(fKArticleTags), fKArticleTags); err != nil {
		o.Rollback()
		return err
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
	orm := orm.NewOrm()

	_, err = orm.Raw(`
		DELETE FROM t_article 
		WHERE id in (?)
	`, ids).Exec()

	return
}
func GetAllArticle(title string, writer string, tag string, state int8, psize int, pinx int) (articles []Article, err error) {
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

	if err != nil {
		return
	}

	for i, _ := range articles {
		v := &articles[i]
		_, err = orm.Raw(`
			select t_tag.* 
				from t_tag 
				inner join t_pk_article_tag on t_tag.id = t_pk_article_tag.tag_id 
			where t_pk_article_tag.article_id = ?
		`, v.Id).QueryRows(&v.Tags)

		if err != nil {
			return
		}
	}

	return
}
func GetHomeArticle(psize, pinx int) (articles []Article, err error) {
	o := orm.NewOrm()

	_, err = o.QueryTable("t_article").Filter("State", 0 /*已发布的*/).OrderBy("-CreateOn").Limit(psize, psize*(pinx-1)).All(&articles)
	if err != nil {
		return
	}

	//	for i, _ := range articles {
	//		v := &articles[i]
	//		_, err = o.Raw(`
	//			select t_tag.*
	//				from t_tag
	//				inner join t_pk_article_tag on t_tag.id = t_pk_article_tag.tag_id
	//			where t_pk_article_tag.article_id = ?
	//		`, v.Id).QueryRows(&v.Tags)

	//		if err != nil {
	//			return
	//		}
	//	}

	return
}
func GetOneArticle(id int) (article *Article, err error) {
	o := orm.NewOrm()

	article = new(Article)
	article.Id = id
	article.CreateOn.Format("yyyy/mm/dd")
	if err = o.Read(article); err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return
}
func GetCountByArticle(state int8) (count int64, err error) {
	o := orm.NewOrm()

	count, err = o.QueryTable("t_article").Filter("state", state).Count()

	return
}

func GetAllTag() (tags []Tag, err error) {
	orm := orm.NewOrm()

	_, err = orm.QueryTable("t_tag").All(&tags)

	return
}
