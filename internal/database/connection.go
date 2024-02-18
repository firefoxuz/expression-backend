package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var (
	dbConnection *sqlx.DB
)

func GetConnection() (*sqlx.DB, error) {
	if dbConnection == nil {
		dbSourceName := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", "postgres", viper.GetString("database.name"), viper.GetString("database.username"), viper.GetString("database.password"))

		db, err := sqlx.Connect("postgres", dbSourceName)
		if err != nil {
			return nil, err
		}

		dbConnection = db
	}

	return dbConnection, nil
}
