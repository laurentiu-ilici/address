package main

import (
	"address/controllers"
	"address/migrations"
	"address/repositories"
	_ "address/routers"

	"github.com/astaxie/beego"
)

func bootstrap() {
	connectionString := beego.AppConfig.String("dbconstring")
	repository := repositories.AddressRepository{connectionString}
	controllers.InitDependencies(repository)
}

func main() {
	migrations.Migrate(beego.AppConfig.String("dbconstring"))

	bootstrap()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
