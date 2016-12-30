package repositories

import (
	"address/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type AddressRepository struct {
	ConnectionString string
}

// FindByPostCode return a list of addresses which are valid for the given post code
func (repo AddressRepository) FindByPostCode(postCode string) ([]models.Address, error) {
	db, err := gorm.Open("postgres", repo.ConnectionString)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var addresses []models.Address

	db.Where("post_code = ?", postCode).Find(&addresses)

	return addresses, nil
}

func (repo AddressRepository) GetAllAddresses() ([]models.Address, error) {
	db, err := gorm.Open("postgres", repo.ConnectionString)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var addresses []models.Address

	db.Find(&addresses)

	return addresses, nil
}
