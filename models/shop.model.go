package models

import (
	"ewarung-api-experiment/db"
	"fmt"
)

type Shop struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetAllShop() (Response, error) {
	var obj Shop
	var arrobj []Shop
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, name, description FROM shops"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.ID,
			&obj.Name,
			&obj.Description)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Message = "Success get all shop"
	res.Data = arrobj

	return res, nil
}

func GetShopById(id int) (Response, error) {
	var obj Shop
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, name, description FROM shops WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.ID,
			&obj.Name,
			&obj.Description)
		if err != nil {
			return res, err
		}
	}

	res.Message = "Success get shop"
	res.Data = obj

	return res, nil

}

func StoreShop(name string, description string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO shops (name, description) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, description)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Message = "Success add shop"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func UpdateShop(id int, name string, description string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE shops SET name = ?, description = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, description, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Message = "Success update shop"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteShop(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM shops WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Message = "Success delete shop"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
