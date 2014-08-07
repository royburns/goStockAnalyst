package main

import (
	// "github.com/Unknwon/gowalker/utils"
	"github.com/astaxie/beego"
	"github.com/royburns/goStockAnalyst/models"
	_ "github.com/royburns/goStockAnalyst/routers"
)

func main() {

	// beego.HttpPort = utils.Cfg.MustInt("beego", "http_port_"+beego.RunMode)
	models.InitDB()
	models.SetCacher()

	beego.Run()
}
