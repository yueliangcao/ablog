package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/yueliangcao/ablog/models"
	"strconv"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) List() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_list.html"

	var (
		state      int8
		count1     int64
		count2     int64
		err        error
		list       []*models.Article
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
		beego.Error(err.Error())
		return
	}
	if count2, err = models.GetCountByArticle(2); err != nil {
		beego.Error(err.Error())
		return
	}

	switch searchtype {
	case "title":
		list, err = models.GetAllArticle(keyword, "", "", state, psize, pinx)
	case "writer":
		list, err = models.GetAllArticle("", keyword, "", state, psize, pinx)
	case "tag":
		list, err = models.GetAllArticle("", "", keyword, state, psize, pinx)
	default:
		list, err = models.GetAllArticle("", "", "", state, psize, pinx)
		searchtype = "title"
	}

	if err != nil {
		beego.Error(err.Error())
	}

	c.Data["state"] = state
	c.Data["count1"] = count1
	c.Data["count2"] = count2
	c.Data["list"] = list
	c.Data["searchtype"] = searchtype
	c.Data["keyword"] = keyword
}

func (c *ArticleController) Add() {
	var err error

	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_add.html"

	if c.Ctx.Request.Method == "GET" {

	} else {
		valid := validation.Validation{}

		var article = new(models.Article)
		err = c.ParseForm(article)
		if err != nil {
			return
		}

		beego.Debug(article.Tags)

		valid.Required(article.Title, "标题").Message("不能为空")
		valid.Required(article.Content, "正文").Message("不能为空")

		if valid.HasErrors() {
			for _, err1 := range valid.Errors {
				beego.Debug("key=%s,msg=%s", err1.Key, err1.Message)
			}

			c.Data["errmsg"] = valid.Errors
			return
		}

		article.Writer = c.User.Usn

		if err = models.AddArticle(article); err != nil {
			beego.Warning(err.Error())
			return
		}
	}
}

func (c *ArticleController) Del() {
	ids1 := c.GetStrings("ids")
	ids := make([]int, 0)

	for _, v := range ids1 {
		if i, _ := strconv.Atoi(v); i > 0 {
			ids = append(ids, i)
		}
	}

	err := models.DeleteArticles(&ids)
	if err != nil {
		beego.Error(err.Error())
		return
	}

	c.Redirect("/admin/article/list", 301)
}

func (c *ArticleController) Edit() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_edit.html"
}

func (c *ArticleController) UpdateAllState() {

}
