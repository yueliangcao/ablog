package controllers

import (
	"ablog/models"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.TplNames = "article/index.tpl"
	c.Data["article"] = models.GetArticle()
}

func (c *ArticleController) Edit() {
	if c.Ctx.Request.Method == "GET" {
		c.TplNames = "article/edit.tpl"
		c.Data["article"] = models.GetArticle()
	} else if c.Ctx.Request.Method == "POST" {

	} else {
	}
}
