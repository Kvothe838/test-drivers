package db

import (
	"database/sql"
	"fmt"

	"github.com/Kvothe838/drivers-api/model"
)

type UsersPersistance interface {
	SaveUser(user model.User) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	UserHasPermission(userId int64, permissionName string) (*bool, error)
}

type UsersDb struct{}

var DefaultUsersPersistance UsersPersistance = &UsersDb{}

func (_ *UsersDb) SaveUser(user model.User) (*model.User, error) {
	result, err := Db.Exec("INSERT INTO User(username, hash, profile_id) VALUES($1, $2, $3) RETURNING id", user.Username, user.Hash, user.Profile.Id)
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

func (_ *UsersDb) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	row := Db.QueryRow("SELECT username, hash, profile_id FROM User WHERE username = $1", username)
	if err := row.Scan(&user.Username, &user.Hash, &user.Profile.Id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Printf("error getting user by username(username = %s): %v\n", username, err)
		return nil, err
	}

	return &user, nil
}

func (_ *UsersDb) UserHasPermission(userId int64, permissionName string) (*bool, error) {
	var res bool
	row := Db.QueryRow(`SELECT true 
							 FROM User u 
							 JOIN Profile pro ON u.profile_id = pro.id
							 JOIN Profiles_Permissions p ON pro.id = p.profile_id
							 JOIN Permission per ON p.permission_id = per.id
							 WHERE u.id = $1 AND per.name = $2`, userId, permissionName)
	if err := row.Scan(&res); err != nil {
		if err == sql.ErrNoRows {
			res = false
			return &res, nil
		}

		return nil, err
	}

	return &res, nil
}
