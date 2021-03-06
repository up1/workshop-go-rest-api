package user

import (
	"database/sql"
)

// Global variable
var DB *sql.DB

func GetAllUsers() (Users, error) {
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users Users

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
