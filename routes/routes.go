package routes

import (
	"net/http"
	"superstore_api/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello leo")
	})

	//Product
	e.GET("api/product", controllers.GetAllProduct)
	e.GET("api/product/category", controllers.GetProductCategoryTotal)
	e.GET("api/product/subcategory", controllers.GetProductSubcategoryTotal)
	e.POST("api/product/add", controllers.AddProduct)
	e.PUT("api/product/update/:id", controllers.UpdateProduct)

	//Order
	e.GET("api/order", controllers.GetAllOrder)
	e.POST("api/order/add", controllers.AddOrder)

	//Customer
	e.GET("api/customer", controllers.GetAllCustomer)
	e.POST("api/customer/add", controllers.AddCustomer)
	e.GET("api/customer/region", controllers.GetTotalCustomerByRegion)
	e.GET("api/customer/state", controllers.GetTotalCustomerByState)

	return e
}
