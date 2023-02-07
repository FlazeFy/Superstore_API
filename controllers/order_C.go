package controllers

import (
	"net/http"
	"strconv"
	"superstore_api/models"

	"github.com/labstack/echo/v4"
)

func GetAllOrder(c echo.Context) error {
	result, err := models.GetAllOrder()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddOrder(c echo.Context) error {
	sales, err := strconv.ParseFloat(c.FormValue("sales"), 64)
	discount, err := strconv.ParseFloat(c.FormValue("discount"), 64)
	profit, err := strconv.ParseFloat(c.FormValue("profit"), 64)
	quantity, err := strconv.ParseInt(c.FormValue("quantity"), 10, 64)

	ShipDate := c.FormValue("ship_date")
	ShipMode := c.FormValue("ship_mode")
	CustomerId := c.FormValue("customer_id")
	ProductId := c.FormValue("product_id")
	Sales := sales
	Quantity := quantity
	Discount := discount
	Profit := profit

	result, err := models.AddOrder(ShipDate, ShipMode, CustomerId, ProductId, Sales, Quantity, Discount, Profit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateOrder(c echo.Context) error {
	sales, err := strconv.ParseFloat(c.FormValue("sales"), 64)
	discount, err := strconv.ParseFloat(c.FormValue("discount"), 64)
	profit, err := strconv.ParseFloat(c.FormValue("profit"), 64)
	quantity, err := strconv.ParseInt(c.FormValue("quantity"), 10, 64)

	OrderId := c.Param("id")
	OrderDate := c.FormValue("order_date")
	ShipDate := c.FormValue("ship_date")
	ShipMode := c.FormValue("ship_mode")
	CustomerId := c.FormValue("customer_id")
	ProductId := c.FormValue("product_id")
	Sales := sales
	Quantity := quantity
	Discount := discount
	Profit := profit

	result, err := models.UpdateOrder(OrderId, OrderDate, ShipDate, ShipMode, CustomerId, ProductId, Sales, Quantity, Discount, Profit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
