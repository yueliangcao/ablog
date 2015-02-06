package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yueliangcao/ablog/logs"
	"github.com/yueliangcao/ablog/models"
	_ "strconv"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Index() {
	c.TplNames = "article/index.tpl"
	id, err := c.GetInt(":id")
	if err != nil {
		logs.Log().Debug("GetInt", err.Error())
	} else {
		logs.Log().Trace("id = %n", id)
	}
	c.Data["article"], err = models.GetOneArticle(id)
	if err != nil {
		logs.Log().Debug("GetOneArticle", err.Error())
	}
}

func (c *ArticleController) Edit() {
	if c.Ctx.Request.Method == "GET" {
		c.TplNames = "article/edit.tpl"
		id, err := c.GetInt("id")
		if err != nil {
			logs.Log().Debug("", err.Error())
		}
		c.Data["article"], err = models.GetOneArticle(id)
		if err != nil {
			logs.Log().Debug("", err.Error())
		}
	} else if c.Ctx.Request.Method == "POST" {

	} else {
	}
}
