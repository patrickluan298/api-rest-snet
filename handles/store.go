package handles

import (
	"api-rest-snet/models"
	"api-rest-snet/repositories"
	"api-rest-snet/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func NewStore(c echo.Context) error {
	request := &models.Store{}
	if err := c.Bind(request); err != nil {
		return err
	}

	if err := services.InsertStore(request); err != nil {
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

func GetStore(c echo.Context) error {
	NumLoja, _ := strconv.Atoi(c.Param("id"))
	err := repositories.SelectStore(NumLoja)

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

func UpdateStore(c echo.Context) error {
	request := &models.Store{}
	if err := c.Bind(request); err != nil {
		return err
	}

	NumLoja, _ := strconv.Atoi(c.Param("id"))

	if err := services.UpdateStore(request, NumLoja); err != nil {
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

func GetAllStores(c echo.Context) error {
	err := repositories.SelectAllStores()

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

func DeleteStore(c echo.Context) error {
	NumLoja, _ := strconv.Atoi(c.Param("id"))
	err := repositories.DeleteStore(NumLoja)

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
