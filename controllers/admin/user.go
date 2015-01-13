package admin

import (
	_ "ablog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
