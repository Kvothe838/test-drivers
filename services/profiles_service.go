package services

import (
	"fmt"

	"github.com/Kvothe838/drivers-api/db"
	"github.com/Kvothe838/drivers-api/model"
)

const DriverProfile = "driver"

func GetProfiles() ([]model.Profile, error) {
	profiles, err := db.DefaultProfilesPersistance.GetProfiles()
	if err != nil {
		fmt.Printf("error getting all profiles: %v\n", err)
		return nil, err
	}

	return profiles, nil
}

func FilterProfile(profiles []model.Profile, profileName string) *model.Profile {
	for _, profile := range profiles {
		if profile.Name == profileName {
			return &profile
		}
	}

	return nil
}
