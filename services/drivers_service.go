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

const driversPerPage = 50

func SaveDriver(newDriver model.Driver) {
	newDriver.User.Profile = profileDriver
	allDrivers = append(allDrivers, newDriver)
}

func GetDrivers(page int) ([]model.Driver, error) {
	return allDrivers[page*50 : (page*50)+50], nil
}
