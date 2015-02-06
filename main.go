package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	_ "github.com/yueliangcao/ablog/routers"
)

var globalSessions *session.Manager

func main() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()

	beego.Run()
}
