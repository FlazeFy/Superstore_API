package models

import (
	"superstore_api/database"
	"net/http"
)

type Customer struct {
	CustomerId string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Segment string `json:"segment"`
	Country string `json:"country"`
	City string `json:"city"`
	State string `json:"state"`
	PostalCode string `json:"postal_code"`
	Region string `json:"region"`
}

func GetAllCustomer() (Response, error){
	var table = "superstore_customer"
	var obj Customer
	var arrobj []Customer
	var res Response

	con := database.CreateCon()

	sqlStatement := "SELECT * FROM " + table

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next(){
		err = rows.Scan(
			&obj.CustomerId,
			&obj.CustomerName,
			&obj.Segment,
			&obj.Country,
			&obj.City,
			&obj.State,
			&obj.PostalCode,
			&obj.Region)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}