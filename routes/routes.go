package routes

import (
	"api-rest-snet/handles"
	"api-rest-snet/models"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func HandleRequest() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		o := models.Message{
			Message: "OK",
		}

		return c.JSON(http.StatusOK, o)
	})

	e.GET("/establishments", handles.GetAllEstablishments)
	e.POST("/establishments", handles.NewEstablishment)
	e.GET("/establishments/:NumEstablishment", handles.GetEstablishment)
	e.PUT("/establishments/:NumEstablishment", handles.UpdateEstablishment)
	e.DELETE("/establishments/:NumEstablishment", handles.DeleteEstablishment)

	e.GET("/stores", handles.GetAllStores)
	e.POST("/stores", handles.NewStore)
	e.GET("/stores/:NumLoja", handles.GetStore)
	e.PUT("/stores/:NumLoja", handles.UpdateStore)
	e.DELETE("/stores/:NumLoja", handles.DeleteStore)

	e.Logger.Fatal(e.Start(":1323"))
}
