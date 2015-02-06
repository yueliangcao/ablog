package admin

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/yueliangcao/ablog/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Index() {

	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/user_index.html"
}

func (c *UserController) Add() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/user_add.html"
}

func (c *UserController) Del() {

}

func (c *UserController) Edit() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/user_edit.html"
}
