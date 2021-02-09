package models

import (
	"ewarung-api-experiment/db"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	IDRole   int    `json:"id_role"`
}

func GetAllUser() (Response, error) {
	var obj User
	var arrobj []User
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, username, email, password, fullname, id_role FROM users"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.ID,
			&obj.Username,
			&obj.Email,
			&obj.Password,
			&obj.FullName,
			&obj.IDRole)
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

func GetUserById(id int) (Response, error) {
	var obj User
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE id = ?"

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
			&obj.Username,
			&obj.Email,
			&obj.Password,
			&obj.FullName,
			&obj.IDRole)
		if err != nil {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil

}

func StoreUser(username string, email string, rawPassword string, fullname string, idRole int) (Response, error) {
	var res Response

	con := db.CreateCon()

	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	sqlStatement := "INSERT INTO users (username, email, password, fullname, id_role) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, email, password, fullname, idRole)
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

func UpdateUser(id int, username string, email string, rawPassword string, fullname string, idRole int) (Response, error) {
	var res Response

	con := db.CreateCon()

	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	sqlStatement := "UPDATE users SET username = ?, email = ?, password = ?, fullname = ?, id_role = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, email, password, fullname, idRole, id)
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

func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM users WHERE id = ?"

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
