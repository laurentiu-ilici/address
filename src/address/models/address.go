package models

import "fmt"

var (
	AddressList []Address
)

func init() {
	address := Address{"81543", "Konradinstr", "Munich"}
	address2 := Address{"80636", "Konradinstr", "Munich"}
	AddressList = make([]Address, 6)
	AddressList[0] = address
	AddressList[1] = address
	AddressList[2] = address
	AddressList[3] = address2
	AddressList[4] = address2
	AddressList[5] = address2
}

type Address struct {
	PostCode string
	Street   string
	City     string
}

func GetAddress(postCode string) (a []Address, err error) {
	result := make([]Address, 0)
	for _, address := range AddressList {
		fmt.Println("shit", address)
		if address.PostCode == postCode {
			result = append(result, address)
		}
	}
	return result, nil
}

func GetAllAddresses() []Address {
	return AddressList
}
