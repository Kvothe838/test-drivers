package db

import (
	"fmt"

	"github.com/Kvothe838/drivers-api/model"
)

func SaveUser(user model.User) (*model.User, error) {
	id, err := Db.Exec("INSERT INTO User(username, hash) VALUES($1, $2) RETURNING id", user.Username, user.Hash)
	if err != nil {
		return nil, fmt.Errorf("error inserting into user: %v", err)
	}

	user.Id, err = id.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting id from saved user: %v", err)
	}

	return &user, nil
}
