package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/yueliangcao/ablog/logs"
	"github.com/yueliangcao/ablog/models"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Index() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_index.html"
}

func (c *ArticleController) Add() {
	user := c.GetSession("user").(*models.User)
	if user == nil {
		c.Redirect("/admin/login", 301)
	}

	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_add.html"

	if c.Ctx.Request.Method == "GET" {

	} else {
		valid := validation.Validation{}

		var article = new(models.Article)
		err := c.ParseForm(article)
		if err != nil {
			return
		}

		valid.Required(article.Title, "title")
		valid.Required(article.Content, "content")

		if valid.HasErrors() {
			for _, err := range valid.Errors {
				logs.Log().Debug("key=%s,msg=%s", err.Key, err.Message)
			}

			return
		}

		article.Writer = user.Usn

		models.AddArticle(article)
	}
}

func (c *ArticleController) Del() {

}

func (c *ArticleController) Edit() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_edit.html"
}
