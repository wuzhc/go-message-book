package routers

import (
	"fmt"
	"my-web/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wuzhc/*.*", &controllers.WuzhcController{})
	beego.Router("/home/get/:id:int/:user:string", &controllers.HomeController{}, "get:Get")
	beego.AutoRouter(&controllers.HomeController{})

	beego.Get("/hello", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})

	ns := beego.NewNamespace("v1",
		// 进入v1的条件,为false时,不会进入
		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "api.beego.me" {
				return false
			} else {
				return true
			}
		}),
		// 进入v1之前做一些东西
		beego.NSBefore(func(ctx *context.Context) {
			fmt.Println("helo world")
		}),
		// 相当于beego.Get
		beego.NSGet("/notallowd", func(ctx *context.Context) {
			ctx.Output.Body([]byte("not allowd"))
		}),
		// 相当于beego.Router
		beego.NSRouter("/hello", &controllers.WuzhcController{}),
		// 嵌套一个namespace
		beego.NSNamespace("/shop",
			beego.NSGet("/:id", func(ctx *context.Context) {
				ctx.Output.Body([]byte("shopinfo"))
			}),
		),
	)
	beego.AddNamespace(ns)
}
