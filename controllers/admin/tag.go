package admin

import (
	"github.com/astaxie/beego"
	_ "github.com/yueliangcao/ablog/models"
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
