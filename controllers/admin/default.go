package controllers

import (
	"ablog/models"
	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

func (c *DefaultController) Index() {
	c.TplNames = "index.tpl"
	c.Data["articles"] = models.GetArticles()
}

func (c *DefaultController) Login() {
	user := new(models.User)
	c.ParseForm(user)

	if user.Usn == "admin" && user.Pwd == "123" {
		c.Data["json"] = map[string]interface{}{"succ": true}
	} else {
		c.Data["json"] = map[string]interface{}{"succ": false, "msg": "帐号或密码错误"}
	}

	c.ServeJson()
}
