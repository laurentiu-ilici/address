package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["address/controllers:AddressController"] = append(beego.GlobalControllerRouter["address/controllers:AddressController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["address/controllers:AddressController"] = append(beego.GlobalControllerRouter["address/controllers:AddressController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:postcode`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
