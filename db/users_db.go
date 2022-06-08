package db

import (
	"database/sql"
	"fmt"

	"github.com/Kvothe838/drivers-api/model"
)

func SaveUser(user model.User) (*model.User, error) {
	result, err := Db.Exec("INSERT INTO User(username, hash) VALUES($1, $2) RETURNING id", user.Username, user.Hash)
	if err != nil {
		fmt.Printf("error inserting into user: %v\n", err)
		return nil, err
	}

	user.Id, err = result.LastInsertId()
	if err != nil {
		fmt.Printf("error getting id from saved user: %v", err)
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	row := Db.QueryRow("SELECT username, hash FROM User WHERE username = $1", username)
	if err := row.Scan(&user.Username, &user.Hash); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Printf("error getting user by username(username = %s): %v\n", username, err)
		return nil, err
	}

	return &user, nil
}
