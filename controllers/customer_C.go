package controllers

import (
	"net/http"
	"superstore_api/models"

	"github.com/labstack/echo/v4"
)

func GetAllCustomer(c echo.Context) error {
	result, err := models.GetAllCustomer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalCustomerByState(c echo.Context) error {
	result, err := models.GetTotalCustomerByState()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalCustomerByRegion(c echo.Context) error {
	result, err := models.GetTotalCustomerByRegion()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddCustomer(c echo.Context) error {
	CustomerName := c.FormValue("customer_name")
	Segment := c.FormValue("segment")
	Country := c.FormValue("country")
	City := c.FormValue("city")
	State := c.FormValue("state")
	PostalCode := c.FormValue("postal_code")
	Region := c.FormValue("region")

	result, err := models.AddCustomer(CustomerName, Segment, Country, City, State, PostalCode, Region)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
