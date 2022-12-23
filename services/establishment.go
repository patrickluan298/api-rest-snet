package services

import (
	"api-rest-snet/models"
	"api-rest-snet/repositories"
	"fmt"

	"gopkg.in/validator.v2"
)

func InsertEstablishment(request *models.Establishment) error {
	if err := validator.Validate(request); err != nil {
		fmt.Println("Error registering establishment:", err.Error())
		return err
	}
	if err := repositories.InsertEstablishment(request); err != nil {
		return err
	}
	return nil
}

func UpdateEstablishment(request *models.Establishment, id int) error {
	if err := validator.Validate(request); err != nil {
		fmt.Println("Error registering establishment:", err.Error())
		return err
	}
	repositories.UpdateEstablishment(request, id)
	return nil
}
