package db

import (
	"database/sql"
	"fmt"

	"github.com/Kvothe838/drivers-api/model"
)

func GetConfiguration() (*model.Configuration, error) {
	var config model.Configuration

	row := Db.QueryRow("SELECT rows_per_page FROM Configuration WHERE")
	if err := row.Scan(&config.RowsPerPage); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error getting configuration: %v", err)
	}

	return &config, nil
}
