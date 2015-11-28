package app

import (
	c "config"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	database *sql.DB
}

func (dbconf *c.DBConfig) NewDatabase() (*Database, error) {
	driverConnectionString := dbconf.DBConnectionString()

	db, err := sql.Open(dbcongf.driver, driver)
	return &Database{
		database: db,
	}, err
}

func