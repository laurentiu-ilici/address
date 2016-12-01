package models

import "errors"

var (
	AddressList map[string]*Address
)

func init() {
	AddressList = make(map[string]*Address)
	u := Address{"81543", "Konradinstr", "Munich"}
	AddressList["81543"] = &u
	AddressList["80636"] = &u
}

type Address struct {
	PostCode string
	Street   string
	City     string
}

func GetAddress(postCode string) (u *Address, err error) {
	if u, ok := AddressList[postCode]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllAddresses() map[string]*Address {
	return AddressList
}
