package handles

import (
	"api-rest-snet/models"
	"api-rest-snet/repositories"
	"api-rest-snet/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func NewEstablishment(c echo.Context) error {
	request := &models.Establishment{}
	if err := c.Bind(request); err != nil {
		return err
	}

	if err := services.InsertEstablishment(request); err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, message)
	} else {
		message := models.Message{
			Message: "OK",
		}
		return c.JSON(http.StatusOK, message)
	}
}

func GetEstablishment(c echo.Context) error {
	NumEstablishment, _ := strconv.Atoi(c.Param("NumEstablishment"))
	err := repositories.SelectEstablishment(NumEstablishment)

	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, message)
	} else {
		message := models.Message{
			Message: "OK",
		}
		return c.JSON(http.StatusOK, message)
	}
}

func UpdateEstablishment(c echo.Context) error {
	request := &models.Establishment{}
	if err := c.Bind(request); err != nil {
		return err
	}

	NumEstablishment, _ := strconv.Atoi(c.Param("NumEstablishment"))

	if err := services.UpdateEstablishment(request, NumEstablishment); err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, message)
	} else {
		message := models.Message{
			Message: "OK",
		}
		return c.JSON(http.StatusOK, message)
	}
}

func GetAllEstablishments(c echo.Context) error {
	err := repositories.SelectAllEstablishments()

	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, message)
	} else {
		message := models.Message{
			Message: "OK",
		}
		return c.JSON(http.StatusOK, message)
	}
}

func DeleteEstablishment(c echo.Context) error {
	NumEstablishment, _ := strconv.Atoi(c.Param("NumEstablishment"))
	err := repositories.DeleteEstablishment(NumEstablishment)

	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, message)
	} else {
		message := models.Message{
			Message: "OK",
		}
		return c.JSON(http.StatusOK, message)
	}
}
