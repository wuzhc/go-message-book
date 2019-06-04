package main

import (
	"fmt"
	_ "message-book/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/down1", "download1")
	beego.BConfig.AppName = "hello world"
	beego.BConfig.RouterCaseSensitive = false // 不区分大小写
	beego.BConfig.Listen.EnableAdmin = true
	beego.BConfig.WebConfig.Session.SessionOn = true
	mysqluser := beego.AppConfig.String("mysqluser")
	fmt.Println(mysqluser)

	beego.SetLogFuncCall(true)
	beego.Run()
}
