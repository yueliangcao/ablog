package main

import (
	"github.com/astaxie/beego"
	_ "github.com/yueliangcao/ablog/routers"
)

func main() {
	beego.SetLogFuncCall(true)
	beego.Run()
	beego.Debug()
}
