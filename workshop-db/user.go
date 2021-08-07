package demo

import "database/sql"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func GetAllUsers(db *sql.DB) (Users, error) {
	rows, err := db.Query("SELECT * FROM users")
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
