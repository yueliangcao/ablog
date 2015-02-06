package logs

import (
	"github.com/astaxie/beego/logs"
)

var log *logs.BeeLogger

func init() {
	log = logs.NewLogger(1000)
	log.SetLogger("console", "")
}

func Log() *logs.BeeLogger {
	return log
}
