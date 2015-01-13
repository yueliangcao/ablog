package models

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
	Name  string `orm:"size(20);index"`
	Count int
}

//文章标签关联
type ArticleTag struct {
	Id        int
	ArticleId int
	TagId     int
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
