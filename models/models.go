package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yueliangcao/ablog/logs"
)

//用户
type User struct {
	Id  int
	Usn string
	Pwd string
}

//文章
type Article struct {
	Id      int
	Title   string
	Content string
	Writer  string
}

//标签
type Tag struct {
	Id    int
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
	db, err = sql.Open("mysql", "root:123456@/ablog?charset=utf8")
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

func GetAllArticle() (articles []Article, err error) {
	var rows *sql.Rows

	dbf(func(db *sql.DB) {
		rows, err = db.Query("select * from t_article")
	})

	if err != nil {
		return nil, err
	}

	articles = make([]Article, 0)
	article := new(Article)
	for rows.Next() {
		err = rows.Scan(&article.Id, &article.Title, &article.Content, &article.Writer)
		if err != nil {
			return
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
		err = rows.Scan(&article.Id, &article.Title, &article.Content, &article.Writer)
		if err != nil {
			return nil, err
		}
	}

	return article, nil
}
