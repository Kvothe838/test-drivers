package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Kvothe838/drivers-api/model"
)

func GetDrivers(page int, rowsPerPage int) ([]model.Driver, error) {
	rows, err := Db.Query(`SELECT d.dni, d.name, d.surname, u.username
						   FROM Driver d
						   JOIN User u
						   LIMIT $1,$2;`, page, rowsPerPage)
	if err != nil {
		fmt.Printf("error getting all drivers: %v\n", err)
		return nil, err
	}

	drivers, err := scanDriverRows(rows)
	if err != nil {
		return nil, err
	}

	return drivers, nil
}

func GetNonTravellingDrivers() ([]model.Driver, error) {
	now := time.Now().Unix()
	rows, err := Db.Query(`SELECT d.dni, d.name, d.surname, u.username
						   FROM Driver d
						   JOIN User u ON d.user_id = u.id
						   LEFT JOIN Travel t ON t.driver_id = u.id
						   WHERE t.id IS NULL OR t.END < $1
						   GROUP BY d.id, u.id`, now)
	if err != nil {
		fmt.Printf("error getting non travelling drivers: %v\n", err)
		return nil, err
	}

	drivers, err := scanDriverRows(rows)
	if err != nil {
		return nil, err
	}

	return drivers, nil
}

func SaveDriver(driver model.Driver) error {
	savedUser, err := SaveUser(driver.User)
	if err != nil {
		fmt.Printf("error saving user at SaveDriver: %v", err)
		return err
	}

	_, err = Db.Exec("INSERT INTO Driver(dni, name, surname, user_id) VALUES($1, $2, $3)", driver.DNI, driver.Name, driver.Surname, savedUser.Id)
	if err != nil {
		fmt.Printf("error inserting into user: %v\n", err)
		return err
	}

	return nil
}

func scanDriverRows(rows *sql.Rows) ([]model.Driver, error) {
	var drivers []model.Driver

	defer rows.Close()

	for rows.Next() {
		var driver model.Driver

		if err := rows.Scan(&driver.DNI, &driver.Name, &driver.Surname, &driver.User.Username); err != nil {
			fmt.Printf("error scanning driver: %v\n", err)
			return nil, err
		}

		drivers = append(drivers, driver)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("error scanning rows when getting drivers: %v\n", err)
		return nil, err
	}

	return drivers, nil
}
