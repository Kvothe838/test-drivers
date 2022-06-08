package services

import (
	"fmt"

	"github.com/Kvothe838/drivers-api/db"
	"github.com/Kvothe838/drivers-api/model"
)

var permissionCreateUsers model.Permission = model.Permission{
	Name: "createUsers",
}
var profileDriver model.Profile = model.Profile{
	Permissions: []model.Permission{permissionCreateUsers},
}

func SaveDriver(newDriver model.Driver) error {
	newDriver.User.Profile = profileDriver

	err := db.SaveDriver(newDriver)
	if err != nil {
		fmt.Printf("error saving driver: %v\n", err)
		return err
	}

	return nil
}

func GetDrivers(page int) ([]model.Driver, error) {
	config, err := db.GetConfiguration()
	if err != nil {
		fmt.Printf("error getting configuration: %v\n", err)
		return nil, err
	}

	drivers, err := db.GetDrivers(page, config.RowsPerPage)
	if err != nil {
		fmt.Printf("error getting drivers: %v\n", err)
		return nil, err
	}

	return drivers, nil
}

func GetNonTravellingDrivers() ([]model.Driver, error) {
	drivers, err := db.GetNonTravellingDrivers()
	if err != nil {
		fmt.Printf("error getting non travelling drivers: %v\n", err)
		return nil, err
	}

	return drivers, nil
}
