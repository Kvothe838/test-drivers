package db

import (
	"fmt"

	"github.com/Kvothe838/drivers-api/model"
)

type ProfilesPersistance interface {
	GetProfiles() ([]model.Profile, error)
}

type ProfilesDb struct{}

var DefaultProfilesPersistance ProfilesPersistance = &ProfilesDb{}

func (_ *ProfilesDb) GetProfiles() ([]model.Profile, error) {
	rows, err := Db.Query(`SELECT id, name
						   FROM Profile`)
	if err != nil {
		fmt.Printf("error getting all profiles: %v\n", err)
		return nil, err
	}

	var profiles []model.Profile

	defer rows.Close()

	for rows.Next() {
		var profile model.Profile

		if err := rows.Scan(&profile.Id, &profile.Name); err != nil {
			fmt.Printf("error scanning profile: %v\n", err)
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("error scanning rows when getting profiles: %v\n", err)
		return nil, err
	}

	return profiles, nil
}
