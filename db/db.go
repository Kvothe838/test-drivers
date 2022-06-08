package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type DbConnection struct {
	*sql.DB
}

var Db DbConnection

func (db *DbConnection) PingOrDie() {
	if err := db.Ping(); err != nil {
		fmt.Printf("can't reach databse, error: %v\n", err)
	}
}

func (db *DbConnection) Ping() error {
	_, err := db.Exec("SELECT 1")
	return err
}

func InitDatabase() {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "drivers_api",
		AllowNativePasswords: true,
	}

	dbConnection, err := sql.Open("mysql", config.FormatDSN())
	defer dbConnection.Close()

	if err != nil {
		fmt.Printf("error opening connection: %v", err)
	}

	Db = DbConnection{dbConnection}
	Db.PingOrDie()
}
