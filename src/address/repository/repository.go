package repository

import (
	"address/models"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// FindByPostCode return a list of addresses which are valid for the given post code
func FindByPostCode(postCode string) []models.Address {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=address sslmode=disable password=password")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	var addresses []models.Address

	db.Where("post_code = ?", postCode).Find(&addresses)

	return addresses
}

func GetAllAddresses() []models.Address {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=address sslmode=disable password=password")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	var addresses []models.Address

	db.Find(&addresses)

	return addresses
}
