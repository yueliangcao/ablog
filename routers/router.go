package routers

import (
	"ablog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "post:Login")
	beego.Router("/article/:id:int", &controllers.ArticleController{})
	beego.Router("/article/del/:id:int", &controllers.ArticleController{}, "get:Del")
	beego.Router("/article/edit/:id:int", &controllers.ArticleController{}, "*:Edit")
}
