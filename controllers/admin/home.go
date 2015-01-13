package admin

import (
	_ "ablog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/home_index.html"
}

func (c *HomeController) Login() {
	c.TplNames = "admin/home_login.html"
}

func (c *HomeController) Setting() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/home_setting.html"
}
