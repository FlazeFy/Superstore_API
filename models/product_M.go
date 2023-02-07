package models

import (
	"net/http"
	"strconv"
	"superstore_api/database"
	"superstore_api/generator"
)

type (
	Product struct {
		ProductId   string `json:"product_id"`
		ProductName string `json:"product_name"`
		Category    string `json:"category"`
		Subcategory string `json:"subcategory"`
	}

	ProductSubcategoryTotal struct {
		Category    string `json:"category"`
		Subcategory string `json:"subcategory"`
		Total       int64  `json:"total"`
	}

	ProductCategoryTotal struct {
		Category string `json:"category"`
		Total    int64  `json:"total"`
	}
)

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
	res.Message = "Successfully collect " + strconv.Itoa(len(arrobj)) + " data"
	res.Data = arrobj

	return res, nil
}

func GetProductCategoryTotal() (Response, error) {
	var table = "superstore_product"
	var obj ProductCategoryTotal
	var arrobj []ProductCategoryTotal
	var res Response

	con := database.CreateCon()

	sqlStatement := "SELECT category, count(1) as total FROM " + table + " group by 1 order by 2 DESC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Category, &obj.Total)
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

func GetProductSubcategoryTotal() (Response, error) {
	var table = "superstore_product"
	var obj ProductSubcategoryTotal
	var arrobj []ProductSubcategoryTotal
	var res Response

	con := database.CreateCon()

	sqlStatement := "SELECT category, subcategory, count(1) as total FROM " + table + " group by 2 order by 1, 3 DESC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Category, &obj.Subcategory, &obj.Total)
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

func UpdateProduct(ProductId string, ProductName string, Category string, Subcategory string) (Response, error) {
	var res Response

	con := database.CreateCon()

	sqlStatement := "UPDATE superstore_product SET product_name = ?, category = ?, subcategory = ? WHERE product_id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(ProductName, Category, Subcategory, ProductId)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
