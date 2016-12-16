package main

import (
	"address/migrations"
	_ "address/routers"

	"github.com/astaxie/beego"
)

func main() {
	migrations.Migrate()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
