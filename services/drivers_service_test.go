package services

import (
	"testing"

	"github.com/Kvothe838/drivers-api/db"
	"github.com/Kvothe838/drivers-api/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

type ConfigurationMock struct{}
type DriversMock struct{}

func (_ *DriversMock) GetDrivers(page int, rowsPerPage int) ([]model.Driver, error) {
	allDrivers := make([]model.Driver, 0)

	for i := 1; i <= 150; i++ {
		allDrivers = append(allDrivers, model.Driver{
			Id: i,
		})
	}

	return paginate(allDrivers, page, rowsPerPage), nil
}

func paginate(drivers []model.Driver, page int, size int) []model.Driver {
	start := page * size
	end := start + size

	if start > len(drivers) {
		start = len(drivers)
	}

	if end > len(drivers) {
		end = len(drivers)
	}

	return drivers[start:end]
}

func (_ *DriversMock) GetNonTravellingDrivers() ([]model.Driver, error) {
	return nil, nil
}

func (_ *DriversMock) SaveDriver(newDriver model.Driver) error {
	return nil
}

func (_ *ConfigurationMock) GetConfiguration() (*model.Configuration, error) {
	return &model.Configuration{
		RowsPerPage: 50,
	}, nil
}

func TestGetDrivers(t *testing.T) {
	db.DefaultConfigurationPersistance = &ConfigurationMock{}
	db.DefaultDriversPersistance = &DriversMock{}

	assert := assert.New(t)

	driversPage0, err := GetDrivers(0)

	assert.NoErrorf(err, "GetDrivers err: %v", err)
	assert.Equal(50, len(driversPage0), "Drivers length should be same as rows per page")
	assert.Equal(1, driversPage0[0].Id, "First driver on page 0 should have id 1")
	assert.Equal(50, driversPage0[49].Id, "Last driver on page 0 should have id 50")

	driversPage1, err := GetDrivers(1)

	assert.NoErrorf(err, "GetDrivers err: %v", err)
	assert.Equal(50, len(driversPage1), "Drivers length should be same as rows per page")
	assert.Equal(51, driversPage1[0].Id, "First driver on page 1 should have id 51")
	assert.Equal(100, driversPage1[49].Id, "Last driver on page 1 should have id 100")

	driversPage2, err := GetDrivers(2)

	assert.NoErrorf(err, "GetDrivers err: %v", err)
	assert.Equal(50, len(driversPage2), "Drivers length should be same as rows per page")
	assert.Equal(101, driversPage2[0].Id, "First driver on page 2 should have id 101")
	assert.Equal(150, driversPage2[49].Id, "Last driver on page 2 should have id 150")
}
