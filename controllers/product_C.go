package controllers

import (
	"net/http"
	"superstore_api/models"

	"github.com/labstack/echo/v4"
)

func GetAllProduct(c echo.Context) error {
	result, err := models.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddProduct(c echo.Context) error {
	ProductName := c.FormValue("product_name")
	Category := c.FormValue("category")
	Subcategory := c.FormValue("subcategory")

	result, err := models.AddProduct(ProductName, Category, Subcategory)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
