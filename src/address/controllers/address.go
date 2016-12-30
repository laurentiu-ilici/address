package controllers

import (
	"address/models"

	"github.com/astaxie/beego"
)

// AddressRepository needed to get addresses
type AddressRepository interface {
	FindByPostCode(postCode string) ([]models.Address, error)
	GetAllAddresses() ([]models.Address, error)
}

// AddressController about Addresses
type AddressController struct {
	beego.Controller
}

var addressRepository AddressRepository

// GetAll addresses
// @Title get all addresses
// @Description get all Addresses
// @Success 200 {object} models.User
// @router / [get]
func (u *AddressController) GetAll() {
	addresses, err := addressRepository.GetAllAddresses()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = addresses
	}
	u.ServeJSON()
}

// Get addresses that have this postcode
// @Title get addresses by post code
// @Description get address by postcode
// @Param postcode
// @Success 200 {object} models.Address
// @Failure 403 :postcode is empty
// @router /:postcode [get]
func (u *AddressController) Get() {
	postcode := u.GetString(":postcode")
	if postcode != "" {
		address, err := addressRepository.FindByPostCode(postcode)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = address
		}
	}
	u.ServeJSON()
}
