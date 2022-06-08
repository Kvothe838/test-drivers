package services

import (
	"fmt"

	"github.com/Kvothe838/drivers-api/db"
	"github.com/Kvothe838/drivers-api/model"
)

func SaveDriver(newDriver model.Driver) error {
	profiles, err := GetProfiles()
	if err != nil {
		fmt.Printf("error getting profiles: %v", err)
		return err
	}

	driverProfile := FilterProfile(profiles, DriverProfile)
	if driverProfile == nil {
		errorDescription := "error filtering driver profile"
		fmt.Println(errorDescription)
		return fmt.Errorf(errorDescription)
	}

	newDriver.User.Profile = *driverProfile

	err = db.SaveDriver(newDriver)
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
