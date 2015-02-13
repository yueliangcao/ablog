package admin

import (
	"github.com/astaxie/beego"
	_ "github.com/yueliangcao/ablog/logs"
	"github.com/yueliangcao/ablog/models"
	"strings"
)

type BaseController struct {
	beego.Controller
	User *models.User
}

func (this *BaseController) Prepare() {
	this.auth()

	this.Layout = "admin/_layout.html"
}

//登录状态验证
func (this *BaseController) auth() {
	controllerName, actionName := this.GetControllerAndAction()
	if user1 := this.GetSession("user"); user1 != nil {
		var ok bool
		this.User, ok = user1.(*models.User)
		if !ok {
			beego.Emergency("auth error")
		}
	}

	if controllerName == "HomeController" && actionName == "Login" {
		if this.User != nil {
			this.Redirect("/admin/", 301)
		}
	} else {
		if this.User == nil {
			arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
			if len(arr) == 2 {
				usn, pwd := arr[0], arr[1]

				user, _ := models.GetOneUserByUsn(usn)
				if user != nil && pwd == models.Md5([]byte(this.getClientIp()+"|"+user.Pwd)) {
					this.User = user
					this.SetSession("user", user)
				}
			}
		}

		if this.User == nil {
			this.Redirect("/admin/login", 301)
		}
	}
}

//获取用户IP地址
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
