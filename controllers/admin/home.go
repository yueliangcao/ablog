package admin

import (
	"github.com/astaxie/beego"
	"github.com/yueliangcao/ablog/logs"
	"github.com/yueliangcao/ablog/models"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/home_index.html"

	if user := c.GetSession("user"); user == nil {
		c.Redirect("/admin/login", 301)
	}
}

func (c *HomeController) Login() {
	c.TplNames = "admin/home_login.html"

	if c.Ctx.Request.Method == "GET" {

	} else {
		usn := c.GetString("usn")
		pwd := c.GetString("pwd")

		user, err := models.GetOneUserByUsn(usn)
		if err != nil {
			logs.Log().Debug("GetOneUserByUsn", err.Error())
			return
		}

		if user == nil {
			c.Data["errmsg"] = "不存在该用户"
		} else if user.Pwd != pwd {
			c.Data["errmsg"] = "密码有误"
		} else {
			c.SetSession("user", user)
			c.Redirect("/admin", 301)
		}
	}

	return
}

func (c *HomeController) Setting() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/home_setting.html"
}
