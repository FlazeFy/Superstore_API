package models

import (
	"net/http"
	"strconv"
	"superstore_api/database"
	"superstore_api/generator"
	"time"
)

type Order struct {
	OrderId    string  `json:"order_id"`
	OrderDate  string  `json:"order_date"`
	ShipDate   string  `json:"ship_date"`
	ShipMode   string  `json:"ship_mode"`
	CustomerId string  `json:"customer_id"`
	ProductId  string  `json:"product_id"`
	Sales      float64 `json:"sales"`
	Quantity   int64   `json:"quantity"`
	Discount   float64 `json:"discount"`
	Profit     float64 `json:"profit"`
}

func GetAllOrder() (Response, error) {
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

	for rows.Next() {
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
	res.Message = "Successfully collect " + strconv.Itoa(len(arrobj)) + " data"
	res.Data = arrobj

	return res, nil
}

func AddOrder(ShipDate string, ShipMode string, CustomerId string, ProductId string, Sales float64, Quantity int64, Discount float64, Profit float64) (Response, error) {
	var res Response
	dt := time.Now()

	order_date := dt.Format("2001-01-01") //Fix this shit
	id := generator.GenerateOrderId(order_date, 2, 6)

	con := database.CreateCon()

	sqlStatement := "INSERT INTO superstore_order(order_id, order_date, ship_date, ship_mode, customer_id, product_id, sales, quantity, discount, profit) VALUES (?,?,?,?,?,?,?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, order_date, ShipDate, ShipMode, CustomerId, ProductId, Sales, Quantity, Discount, Profit)
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
