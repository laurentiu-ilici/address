package controllers

import (
	"address/models"
	"address/repository"

	"github.com/astaxie/beego"
)

// AddressController about Users
type AddressController struct {
	beego.Controller
}

// @GetAll addresses
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *AddressController) GetAll() {
	addresses := getAllAddresses()
	u.Data["json"] = addresses
	u.ServeJSON()
}

// @Title Get addresses that have this postcode
// @Description get address by postcode
// @Param	postcode		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :postcode is empty
// @router /:postcode [get]
func (u *AddressController) Get() {
	postcode := u.GetString(":postcode")
	if postcode != "" {
		address, err := getAddress(postcode)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = address
		}
	}
	u.ServeJSON()
}

func getAddress(postCode string) (a []models.Address, err error) {
	return repository.FindByPostCode(postCode), nil
}

func getAllAddresses() []models.Address {
	return repository.GetAllAddresses()
}
