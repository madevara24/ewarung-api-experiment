package models

import (
	"ewarung-api-experiment/db"
	"fmt"
	"log"
)

type Credentials struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetUserByUsername(username string) (Credentials, error) {
	var obj Credentials
	var hasResult bool = false

	con := db.CreateCon()

	sqlStatement := "SELECT u.id, u.username, u.email, u.password, r.name AS role FROM users u JOIN roles r ON u.id_role = r.id WHERE u.username = ?"

	rows, err := con.Query(sqlStatement, username)
	if err != nil {
		fmt.Println(err)
		return obj, err
	}
	defer rows.Close()

	for rows.Next() {
		hasResult = true
		err := rows.Scan(
			&obj.ID,
			&obj.Username,
			&obj.Email,
			&obj.Password,
			&obj.Role)
		if err != nil {
			log.Println(err)
			return obj, err
		}
	}

	if !hasResult {
		fmt.Println("No rows")
		return obj, err
	}
	fmt.Println(obj)
	return obj, nil
}
