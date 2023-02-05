package controllers

import (
	"github.com/labstack/echo/v4"
	"superstore_api/models"
	"net/http"
)

func GetAllCustomer(c echo.Context) error{
	result, err := models.GetAllCustomer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg":err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}