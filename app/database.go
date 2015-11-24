package app

import (
	c "config"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	database *sql.DB
}

func (dbconf *c.DBConfig) NewDatabase() (*Database, error) {
	driverConnectionString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s", dbconf.HostName(), dbconf.DBName(), dbconf.UserName(), dbconf.Password())

	db, err := sql.Open(dbcongf.driver, driver)
	return &Database{
		database: db,
	}, err
}
