package database

import (
	"database/sql"
	"devbook-api/src/infra/config"

	// CONNECTION DRIVER TO MYSQL
	_ "github.com/go-sql-driver/mysql"
)

func GetDbConnection() (*sql.DB, error) {
	stringConnection := config.DatabaseConnectionString

	db, err := sql.Open("mysql", stringConnection)
	return handleErrors(db, err)
}

func handleErrors(db *sql.DB, err error) (*sql.DB, error) {
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
