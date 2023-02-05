package routes

import (
	"github.com/labstack/echo/v4"
	"superstore_api/controllers"
	"net/http"
)

func Init() *echo.Echo  {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello leo")
	})

	//Product
	e.GET("/product", controllers.GetAllProduct)

	return e
}