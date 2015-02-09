package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/yueliangcao/ablog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var err error
	c.TplNames = "index.tpl"
	c.Data["articles"], err = models.GetHomeArticle(100, 1)
	fmt.Print(c.Data["articles"])

	if err != nil {
		beego.Warning(err.Error())
	}
}

func (c *MainController) Login() {
	user := new(models.User)
	c.ParseForm(user)

	if user.Usn == "admin" && user.Pwd == "123" {
		c.Data["json"] = map[string]interface{}{"succ": true}
	} else {
		c.Data["json"] = map[string]interface{}{"succ": false, "msg": "帐号或密码错误"}
	}

	c.ServeJson()
}
