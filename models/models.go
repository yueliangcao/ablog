package models

type User struct {
	Usn string
	Pwd string
}

type Article struct {
	Id      int
	Title   string
	Content string
	Writer  string
}

var articles = []Article{
	Article{1, "我的第一篇日志", "日志的正文", "yueliang"},
	Article{2, "我的第二篇日志", "日志的正文", "yueliang"}}

func GetArticles() []Article {
	return articles
}

func GetArticle() Article {
	return articles[0]
}

func UpdateArticle() error {

}
