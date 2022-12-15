package services

import (
	"api-rest-snet/models"
	"fmt"

	"gopkg.in/validator.v2"
)

func ValidateStore(e *models.Store) error {
	if err := validator.Validate(e); err != nil {
		fmt.Println("Error registering store:", err.Error())
		return err
	}
	return nil
}
