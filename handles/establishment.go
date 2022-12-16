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
	id, _ := strconv.Atoi(c.Param("id"))
	err := repositories.SelectEstablishment(id)

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

	id, _ := strconv.Atoi(c.Param("id"))

	if err := services.UpdateEstablishment(request, id); err != nil {
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
	id, _ := strconv.Atoi(c.Param("id"))
	err := repositories.DeleteEstablishment(id)

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
