package routers

import (
	"github.com/astaxie/beego"
	"github.com/yueliangcao/ablog/controllers"
	"github.com/yueliangcao/ablog/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/login", &controllers.MainController{}, "post:Login")
	beego.Router("/article/:id:int", &controllers.ArticleController{}, "*:Index")
	// beego.Router("/article/del/:id:int", &controllers.ArticleController{}, "get:Del")
	// beego.Router("/article/edit/:id:int", &controllers.ArticleController{}, "*:Edit")
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
