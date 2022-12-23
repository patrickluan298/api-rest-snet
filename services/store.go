package services

import (
	"api-rest-snet/models"
	"api-rest-snet/repositories"
	"fmt"

	"gopkg.in/validator.v2"
)

func InsertStore(request *models.Store) error {
	if err := validator.Validate(request); err != nil {
		fmt.Println("Error registering store:", err.Error())
		return err
	}
	if err := repositories.InsertStore(request); err != nil {
		return err
	}
	return nil
}

func UpdateStore(request *models.Store, id int) error {
	if err := validator.Validate(request); err != nil {
		fmt.Println("Error registering store:", err.Error())
		return err
	}
	repositories.UpdateStore(request, id)
	return nil
}
