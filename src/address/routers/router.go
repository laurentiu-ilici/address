package routers

import (
	"address/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/addresses",
			beego.NSInclude(
				&controllers.AddressController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
