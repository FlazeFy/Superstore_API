package models

import (
	"net/http"
	"superstore_api/database"
	"superstore_api/generator"
)

type Product struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
}

func GetAllProduct() (Response, error) {
	var table = "superstore_product"
	var obj Product
	var arrobj []Product
	var res Response

	con := database.CreateCon()

	sqlStatement := "SELECT * FROM " + table

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
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

func AddProduct(ProductName string, Category string, Subcategory string) (Response, error) {
	var res Response

	id := generator.GenerateProductId(Category, Subcategory, 8)

	con := database.CreateCon()

	sqlStatement := "INSERT INTO superstore_product(product_id, product_name, category, subcategory) VALUES (?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, ProductName, Category, Subcategory)
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
