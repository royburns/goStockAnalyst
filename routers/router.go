package routers

import (
	"github.com/astaxie/beego"
	// "github.com/royburns/goStockAnalyst/apis"
	"github.com/royburns/goStockAnalyst/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/api/getstock", &controllers.IndexController{}, "get:Stock")

	beego.Router("/test", &controllers.TestController{})
	beego.Router("/about", &controllers.AboutController{})

	// static file
	beego.SetStaticPath("/data", "data")
}
