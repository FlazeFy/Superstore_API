package models

import (
	"net/http"
	"superstore_api/database"
	"superstore_api/generator"
)

type Customer struct {
	CustomerId   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Segment      string `json:"segment"`
	Country      string `json:"country"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postal_code"`
	Region       string `json:"region"`
}

func GetAllCustomer() (Response, error) {
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

	for rows.Next() {
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

func AddCustomer(CustomerName string, Segment string, Country string, City string, State string, PostalCode string, Region string) (Response, error) {
	var res Response

	id := generator.GenerateCustomerId(CustomerName, 5)

	con := database.CreateCon()

	sqlStatement := "INSERT INTO superstore_customer(customer_id, customer_name, segment, country, city, state, postal_code, region) VALUES (?,?,?,?,?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, CustomerName, Segment, Country, City, State, PostalCode, Region)
	if err != nil {
		return res, err
	}

	result.LastInsertId()
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"Last Inserted Id": id,
	}
	return res, nil
}
