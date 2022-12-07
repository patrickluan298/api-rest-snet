package handles

import (
	"api-rest-snet/repository"
	"api-rest-snet/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func NewEstablishment(c echo.Context) error {
	r := &models.Establishment{}
	if err := c.Bind(r); err != nil {
		return err
	}
	err := repository.InsertEstablishment(r)

	if err != nil {
		e := models.Message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, e)
	} else {
		o := models.Message{
			Message: "OK",
		}
		return c.JSON(http.StatusOK, o)
	}
}

func GetEstablishment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repository.SelectEstablishment(id)

	if err != nil {
		e := models.Message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, e)
	} else {
		o := models.Message{
			Message: "OK",
		}
		return c.JSON(http.StatusOK, o)
	}
}

func UpdateEstablishment(c echo.Context) error {
	r := &models.Establishment{}
	if err := c.Bind(r); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	repository.UpdateEstablishment(r, id)

	o := models.Message{
		Message: "OK",
	}

	return c.JSON(http.StatusOK, o)
}

func GetAllEstablishments(c echo.Context) error {
	err := repository.SelectAllEstablishments()

	if err != nil {
		e := models.Message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, e)
	} else {
		o := models.Message{
			Message: "OK",
		}
		return c.JSON(http.StatusOK, o)
	}
}

func DeleteEstablishment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repository.DeleteEstablishment(id)

	if err != nil {
		e := models.Message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, e)
	} else {
		o := models.Message{
			Message: "OK",
		}
		return c.JSON(http.StatusOK, o)
	}
}
