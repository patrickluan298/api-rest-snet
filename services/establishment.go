package services

import (
	"api-rest-snet/models"
	"fmt"

	"gopkg.in/validator.v2"
)

func ValidateEstablishment(e *models.Establishment) error {
	if err := validator.Validate(e); err != nil {
		fmt.Println("Error registering establishment:", err.Error())
		return err
	}
	return nil
}

func ValidateStore(e *models.Store) error {
	if err := validator.Validate(e); err != nil {
		fmt.Println("Error registering store:", err.Error())
		return err
	}
	return nil
}
