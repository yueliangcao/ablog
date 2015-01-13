package admin

import (
	_ "ablog/models"
	"github.com/astaxie/beego"
)

type TagController struct {
	beego.Controller
}

func (c *TagController) Index() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/tag_index.html"
}

func (c *TagController) Del() {

}
