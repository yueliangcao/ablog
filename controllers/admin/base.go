package admin

import (
	"github.com/astaxie/beego"
	"github.com/yueliangcao/ablog/models"
)

type BaseController struct {
	beego.Controller
	User *models.User
}

func (this *BaseController) Prepare() {
	if tmp := this.GetSession("user"); tmp == nil {
		this.Redirect("/admin/login", 301)
	} else {
		this.User = tmp.(*models.User)
	}
}
