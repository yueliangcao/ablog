package admin

import (
	"github.com/astaxie/beego"
	"github.com/yueliangcao/ablog/models"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	c.TplNames = "admin/home_index.html"
}

//后台登陆
func (c *HomeController) Login() {
	c.Layout = ""
	c.TplNames = "admin/home_login.html"

	if c.Ctx.Request.Method == "GET" {

	} else {
		usn := c.GetString("usn")
		pwd := c.GetString("pwd")
		remember, err := c.GetBool("remember")

		user, err := models.GetOneUserByUsn(usn)
		if err != nil {
			beego.Error(err.Error())
			return
		}

		if user == nil {
			c.Data["errmsg"] = "不存在该用户"
		} else if user.Pwd != pwd {
			c.Data["errmsg"] = "密码有误"
		} else { //登陆成功
			authkey := models.Md5([]byte(c.getClientIp() + "|" + user.Pwd))
			if remember {
				c.Ctx.SetCookie("auth", usn+"|"+authkey, 7*86400)
			}

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
