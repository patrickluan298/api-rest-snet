package handles

import (
	"api-rest-snet/models"
	"api-rest-snet/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func NewStore(c echo.Context) error {
	r := &models.Store{}
	if err := c.Bind(r); err != nil {
		return err
	}
	err := services.InsertStore(r)

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

func GetStore(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.SelectStore(id)

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

func UpdateStore(c echo.Context) error {
	r := &models.Store{}
	if err := c.Bind(r); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	services.UpdateStore(r, id)

	o := models.Message{
		Message: "OK",
	}

	return c.JSON(http.StatusOK, o)
}

func GetAllStores(c echo.Context) error {
	err := services.SelectAllStores()

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

func DeleteStore(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteStore(id)

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
