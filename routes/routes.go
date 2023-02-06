package routes

import (
	"net/http"
	"superstore_api/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello leo")
	})

	//Product
	e.GET("/product", controllers.GetAllProduct)
	e.GET("/product/category", controllers.GetProductCategoryTotal)
	e.GET("/product/subcategory", controllers.GetProductSubcategoryTotal)
	e.POST("/product/add", controllers.AddProduct)

	//Order
	e.GET("/order", controllers.GetAllOrder)
	e.POST("/order/add", controllers.AddOrder)

	//Customer
	e.GET("/customer", controllers.GetAllCustomer)
	e.POST("/customer/add", controllers.AddCustomer)

	return e
}
