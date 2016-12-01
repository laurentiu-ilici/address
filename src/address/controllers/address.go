package controllers

import (
	"address/models"

	"github.com/astaxie/beego"
)

// AddressController about Users
type AddressController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *AddressController) GetAll() {
	addresses := models.GetAllAddresses()
	u.Data["json"] = addresses
	u.ServeJSON()
}

// @Title Get
// @Description get address by postcode
// @Param	postcode		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :postcode is empty
// @router /:postcode [get]
func (u *AddressController) Get() {
	postcode := u.GetString(":postcode")
	if postcode != "" {
		address, err := models.GetAddress(postcode)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = address
		}
	}
	u.ServeJSON()
}
