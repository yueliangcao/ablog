package models

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yueliangcao/ablog/logs"
	"strings"
	"time"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:!@#$%^&*(0)@tcp(106.187.54.95:3306)/ablog?charset=utf8&parseTime=true&loc=Local")
	orm.RegisterModelWithPrefix("t_", new(User), new(Article), new(Tag), new(FKArticleTag))
	orm.Debug = true
}

const (
	Publish = iota
	Draft
	Recycle
)

//用户
type User struct {
	Id  int
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
	CreateOn time.Time
	UpdateOn time.Time

	Tags []Tag `orm:"-"`
}

//标签
type Tag struct {
	Id    int `orm:"pk"`
	Name  string
	Count int
}

//文章标签关联
type FKArticleTag struct {
	Id        int
	ArticleId int
	TagId     int
}

func openDb() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:!@#$%^&*(0)@tcp(106.187.54.95:3306)/ablog?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		logs.Log().Warning("OpenDB", err.Error())
	}
	return
}

func dbf(f func(db *sql.DB)) {
	db, _ := openDb()
	defer db.Close()

	f(db)
}

func GetOneUserByUsn(usn string) (*User, error) {
	db, err := openDb()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select * from t_user where usn = ? limit 1", usn)
	if err != nil {
		return nil, err
	}

	var user *User = nil
	if rows.Next() {
		user = new(User)
		err = rows.Scan(&user.Id, &user.Usn, &user.Pwd)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func AddArticle(article *Article) (err error) {
	article.CreateOn = time.Now()
	article.UpdateOn = article.CreateOn

	dbf(func(db *sql.DB) {
		_, err = db.Exec("insert into t_article(title,content,writer,top,state,url_name,create_on,update_on) values(?,?,?,?,?,?,?,?)",
			&article.Title, &article.Content, &article.Writer, &article.Top, &article.State, &article.UrlName, &article.CreateOn, &article.UpdateOn)
	})

	return
}
func UpdateArticle(article *Article) (err error) {
	dbf(func(db *sql.DB) {
		_, err = db.Exec("update t_article set title = ?,content = ?,writer = ? where id = ?", &article.Title, &article.Content, &article.Writer, &article.Id)
	})

	return
}
func GetAllArticle(title string, writer string, tag string, state int8, psize int, pinx int) (articles []Article, err error) {
	orm := orm.NewOrm()
	qs := orm.QueryTable("t_article").Filter("state", state)

	if title = strings.TrimSpace(title); title != "" {
		qs = qs.Filter("title__icontains", title)
	}
	if writer = strings.TrimSpace(writer); writer != "" {
		qs = qs.Filter("writer__icontains", writer)
	}

	_, err = qs.Limit(psize, psize*(pinx-1)).All(&articles)
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
	var rows *sql.Rows

	dbf(func(db *sql.DB) {
		rows, err = db.Query("select * from t_article where state = ? order by create_on desc limit ?,?",
			0 /*已发布的*/, psize*(pinx-1), psize)
	})

	if err != nil {
		return nil, err
	}

	articles = make([]Article, 0)
	article := new(Article)
	for rows.Next() {
		err = rows.Scan(
			&article.Id,
			&article.Title,
			&article.Content,
			&article.Writer,
			&article.Top,
			&article.State,
			&article.UrlName,
			&article.Pv,
			&article.CreateOn,
			&article.UpdateOn)
		if err != nil {
			return
		}

		var rows1 *sql.Rows
		dbf(func(db *sql.DB) {
			rows1, err = db.Query(`
				select t_tag.* 
					from t_tag 
					inner join t_pk_article_tag on t_tag.id = t_pk_article_tag.tag_id 
				where t_pk_article_tag.article_id = ?
			`, article.Id)
		})
		if err != nil {
			return
		}

		tag := new(Tag)
		for rows1.Next() {
			err = rows1.Scan(&tag.Id, &tag.Name, &tag.Count)
			if err != nil {
				return
			}
			article.Tags = append(article.Tags, *tag)
		}

		articles = append(articles, *article)
	}

	return
}

func GetOneArticle(id int) (*Article, error) {
	var rows *sql.Rows
	var err error

	dbf(func(db *sql.DB) {
		rows, err = db.Query("select * from t_article where id = ? limit 1", id)
	})

	if err != nil {
		return nil, err
	}

	var article *Article = nil
	if rows.Next() {
		article = new(Article)
		err = rows.Scan(
			&article.Id,
			&article.Title,
			&article.Content,
			&article.Writer,
			&article.Top,
			&article.State,
			&article.UrlName,
			&article.Pv,
			&article.CreateOn,
			&article.UpdateOn)
		if err != nil {
			return nil, err
		}
	}

	return article, nil
}

func GetCountByArticle(state int8) (count int, err error) {
	var rows *sql.Rows

	dbf(func(db *sql.DB) {
		rows, err = db.Query("select count(1) from t_article where state = ?", state)
	})

	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}
