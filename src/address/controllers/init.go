package controllers

var addressRepository AddressRepository

func InitDependencies(repository AddressRepository) {
	addressRepository = repository
}
