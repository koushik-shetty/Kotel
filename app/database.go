package app

import (
	"database/sql"
	"fmt"
	_ "log"

	_ "github.com/lib/pq"

	c "github.com/koushik-shetty/Kotel/config"
)

type Database struct {
	database *sql.DB
}

func NewDatabase(dbconf *c.DBConfig) (*Database, error) {
	db, err := sql.Open(dbconf.Driver(), dbconf.DBConnectionString())
	if err != nil {
		fmt.Errorf("Error connecting to db: %v", err)
		return &Database{
			nil,
		}, err
	}
	return &Database{
		database: db,
	}, err
}

func (db *Database) Database() *sql.DB {
	return db.database
}
