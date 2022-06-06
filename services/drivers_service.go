package services

import (
	"github.com/Kvothe838/drivers-api/model"
)

var permissionCreateUsers model.Permission = model.Permission{
	Name: "createUsers",
}
var profileDriver model.Profile = model.Profile{
	Permissions: []model.Permission{permissionCreateUsers},
}
var allDrivers []model.Driver = make([]model.Driver, 0)
var allTravels []model.Travel = make([]model.Travel, 0)

const driversPerPage = 50

func SaveDriver(newDriver model.Driver) {
	newDriver.User.Profile = profileDriver
	allDrivers = append(allDrivers, newDriver)
}

func GetDrivers(page int) ([]model.Driver, error) {
	return allDrivers[page*50 : (page*50)+50], nil
}

func GetNonTravellingDrivers() ([]model.Driver, error) {
	nonTravellingDrivers := make([]model.Driver, 0)
	isDriverWithTravelByDriverId := make(map[int]bool)

	for _, travel := range allTravels {
		isDriverWithTravelByDriverId[travel.Id] = true
	}

	for _, driver := range allDrivers {
		if !isDriverWithTravelByDriverId[driver.Id] {
			nonTravellingDrivers = append(nonTravellingDrivers, driver)
		}
	}

	return nonTravellingDrivers, nil
}
