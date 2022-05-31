package services

import (
	"fmt"

	"github.com/Kvothe838/drivers-api/model"
)

var permissionCreateUsers model.Permission = model.Permission{
	Name: "createUsers",
}
var profileDriver model.Profile = model.Profile{
	Permissions: []model.Permission{permissionCreateUsers},
}
var allDrivers []model.Driver = make([]model.Driver, 0)

func SaveDriver(newDriver model.Driver) {
	newDriver.User.Profile = profileDriver
	allDrivers = append(allDrivers, newDriver)

	fmt.Printf("new driver: %v\n", newDriver)
	fmt.Printf("all drivers: %v\n", allDrivers)
}
