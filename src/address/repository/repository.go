package repository

import (
	"address/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// FindByPostCode return a list of addresses which are valid for the given post code
func FindByPostCode(postCode string) ([]models.Address, error) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=address sslmode=disable password=password")
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var addresses []models.Address

	db.Where("post_code = ?", postCode).Find(&addresses)

	return addresses, nil
}

func GetAllAddresses() ([]models.Address, error) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=address sslmode=disable password=password")
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var addresses []models.Address

	db.Find(&addresses)

	return addresses, nil
}
