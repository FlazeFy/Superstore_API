package models

import (
	"superstore_api/database"
	"net/http"
)

type Order struct {
	OrderId string `json:"order_id"`
	OrderDate string `json:"order_date"`
	ShipDate string `json:"ship_date"`
	ShipMode string `json:"ship_mode"`
	CustomerId string `json:"customer_id"`
	ProductId string `json:"product_id"`
	Sales float64 `json:"sales"`
	Quantity string `json:"quantity"`
	Discount float64 `json:"discount"`
	Profit float64 `json:"profit"`
}

func GetAllOrder() (Response, error){
	var table = "superstore_order"
	var obj Order
	var arrobj []Order
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
			&obj.OrderId,
			&obj.OrderDate,
			&obj.ShipDate,
			&obj.ShipMode,
			&obj.CustomerId,
			&obj.ProductId,
			&obj.Sales,
			&obj.Quantity,
			&obj.Discount,
			&obj.Profit)

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