package models

import (
	"ewarung-api-experiment/db"
	"fmt"

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

type UserRole struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Role     Role   `json:"role"`
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

	res.Message = "Success get all user"
	res.Data = arrobj

	return res, nil
}

func GetUserById(id int) (Response, error) {
	var obj User
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, username, email, password, fullname, id_role FROM users WHERE id = ?"

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

	res.Message = "Success get user"
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

	res.Message = "Success add user"
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

	res.Message = "Success update user"
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

	res.Message = "Success delete user"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func GetAllUserWithRole() (Response, error) {
	var obj UserRole
	var arrobj []UserRole
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT u.id, u.username, u.email, u.password, u.fullname, u.id_role, r.name FROM users u JOIN roles r on u.id_role = r.id"

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
			&obj.Role.ID,
			&obj.Role.Name)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Message = "Success get all user with roles"
	res.Data = arrobj

	return res, nil
}

func GetUserWithRoleById(id int) (Response, error) {
	var obj UserRole
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT u.id, u.username, u.email, u.password, u.fullname, u.id_role, r.name FROM users u JOIN roles r on u.id_role = r.id WHERE u.id = ?"

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
			&obj.Role.ID,
			&obj.Role.Name)
		if err != nil {
			return res, err
		}
	}

	res.Message = "Success get user with role"
	res.Data = obj

	return res, nil

}
