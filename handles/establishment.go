package handles

import (
	"api-rest-snet/models"
	"api-rest-snet/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func NewEstablishment(c echo.Context) error {
	r := &models.Establishment{}
	if err := c.Bind(r); err != nil {
		return err
	}
	err := repositories.InsertEstablishment(r)

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
	err := repositories.SelectEstablishment(id)

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
	repositories.UpdateEstablishment(r, id)

	o := models.Message{
		Message: "OK",
	}

	return c.JSON(http.StatusOK, o)
}

func GetAllEstablishments(c echo.Context) error {
	err := repositories.SelectAllEstablishments()

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
	err := repositories.DeleteEstablishment(id)

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
