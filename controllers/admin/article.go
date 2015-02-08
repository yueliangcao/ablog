package admin

import (
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/yueliangcao/ablog/logs"
	"github.com/yueliangcao/ablog/models"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) List() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_list.html"

	var (
		state      int8
		count1     int
		count2     int
		err        error
		list       []models.Article
		searchtype string
		keyword    string
		psize      int
		pinx       int
	)

	state, _ = c.GetInt8("state")
	keyword = c.GetString("keyword")
	searchtype = c.GetString("searchtype")
	psize, _ = c.GetInt("psize")
	pinx, _ = c.GetInt("pinx")

	if psize < 1 {
		psize = 10
	}
	if pinx < 1 {
		pinx = 1
	}

	if count1, err = models.GetCountByArticle(1); err != nil {
		logs.Log().Debug("get count1 %s", err.Error())
	}
	if count2, err = models.GetCountByArticle(2); err != nil {
		logs.Log().Debug("get count2 %s", err.Error())
	}

	switch searchtype {
	case "title":
		if list, err = models.GetAllArticle(keyword, "", "", psize, pinx); err != nil {
			logs.Log().Debug("get list %s", err.Error())
		}
	case "tag":
	case "writer":
		if list, err = models.GetAllArticle("", keyword, "", psize, pinx); err != nil {
			logs.Log().Debug("get list %s", err.Error())
		}
	default:
		searchtype = "title"
	}

	c.Data["state"] = state
	c.Data["count1"] = count1
	c.Data["count2"] = count2
	c.Data["list"] = list
	c.Data["searchtype"] = searchtype
	c.Data["keyword"] = keyword
}

func (c *ArticleController) Add() {
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

		valid.Required(article.Title, "标题").Message("不能为空")
		valid.Required(article.Content, "正文").Message("不能为空")

		if valid.HasErrors() {
			for _, err := range valid.Errors {
				logs.Log().Debug("key=%s,msg=%s", err.Key, err.Message)
			}

			c.Data["errmsg"] = valid.Errors
			return
		}

		article.Writer = c.User.Usn

		models.AddArticle(article)
	}
}

func (c *ArticleController) Del() {

}

func (c *ArticleController) Edit() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_edit.html"
}
