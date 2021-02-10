package models

import (
	"ewarung-api-experiment/db"
	"fmt"
	"net/http"
)

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllRole() (Response, error) {
	var obj Role
	var arrobj []Role
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, name FROM roles"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.ID,
			&obj.Name)
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

func StoreRole(name string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO roles (name) VALUES (?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name)
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

func UpdateRole(id int, name string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE roles SET name = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, id)
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

func DeleteRole(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM roles WHERE id = ?"

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

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
