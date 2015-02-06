package admin

import (
	"github.com/astaxie/beego"
	_ "github.com/yueliangcao/ablog/models"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Index() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_index.html"
}

func (c *ArticleController) Add() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_add.html"
}

func (c *ArticleController) Del() {

}

func (c *ArticleController) Edit() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/article_edit.html"
}
