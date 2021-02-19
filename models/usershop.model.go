package models

import (
	"ewarung-api-experiment/db"
	"fmt"
)

type UserShop struct {
	ID   int      `json:"id"`
	User UserRole `json:"user"`
	Shop Shop     `json:"shop"`
}

func GetAllUserShop() (Response, error) {
	var obj UserShop
	var arrobj []UserShop
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT us.id, u.id, u.username, u.email, u.password, u.fullname, u.id_role, r.name, s.id, s.name, s.description FROM users_shops us JOIN users u on us.id_user = u.id JOIN roles r on r.id = u.id_role JOIN shops s on s.id = us.id_shop"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&obj.ID,
			&obj.User.ID,
			&obj.User.Username,
			&obj.User.Email,
			&obj.User.Password,
			&obj.User.FullName,
			&obj.User.Role.ID,
			&obj.User.Role.Name,
			&obj.Shop.ID,
			&obj.Shop.Name,
			&obj.Shop.Description)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Message = "Success get all shop"
	res.Data = arrobj

	return res, nil
}

func GetUserShopById(id int) (Response, error) {
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

func StoreUserShop(id_user int, id_shop int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO users_shops (id_user, id_shop) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id_user, id_shop)
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

func UpdateUserShop(id int, name string, description string) (Response, error) {
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

func DeleteUserShop(id int) (Response, error) {
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
