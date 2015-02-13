package admin

import (
	_ "github.com/astaxie/beego"
	"github.com/yueliangcao/ablog/logs"
	"github.com/yueliangcao/ablog/models"
)

type TagController struct {
	BaseController
}

func (c *TagController) List() {
	c.Layout = "admin/_layout.html"
	c.TplNames = "admin/tag_list.html"

	var (
		list []models.Tag
	)

	list, err := models.GetAllTag()
	if err != nil {
		logs.Log().Debug("getAllTag err %s", err.Error())
	}

	c.Data["list"] = list
}

func (c *TagController) Del() {

}
