package models

import (
	"superstore_api/database"
	"net/http"
)

var table = "superstore_product"

type Product struct {
	ProductId string `json:"product_id"`
	ProductName string `json:"product_name"`
	Category string `json:"category"`
	Subcategory string `json:"subcategory"`
}

func GetAllProduct() (Response, error){
	var obj Product
	var arrobj []Product
	var res Response

	con := database.CreateCon()

	sqlStatement := "SELECT * FROM " + table

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.ProductId, &obj.ProductName, &obj.Category, &obj.Subcategory)
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