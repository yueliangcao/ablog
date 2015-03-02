package routers

import (
	"github.com/astaxie/beego"
	"github.com/yueliangcao/ablog/controllers"
	"github.com/yueliangcao/ablog/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/article/:id:int", &controllers.HomeController{}, "*:Article")
	beego.Router("/tag/:name", &controllers.HomeController{}, "*:Tag")
	beego.Router("/archives", &controllers.HomeController{}, "*:Archives")
	beego.Router("/tags", &controllers.HomeController{}, "*:Tags")
	beego.Router("/aboutme", &controllers.HomeController{}, "*:AboutMe")

	beego.Router("/admin", &admin.HomeController{}, "*:Index")
	beego.Router("/admin/login", &admin.HomeController{}, "*:Login")

	beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/del", &admin.ArticleController{}, "*:Del")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")

	beego.Router("/admin/user", &admin.UserController{}, "*:Index")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/del", &admin.UserController{}, "*:Del")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")

	beego.Router("/admin/tag/list", &admin.TagController{}, "*:List")
	beego.Router("/admin/tag/del", &admin.TagController{}, "*:Del")
}
