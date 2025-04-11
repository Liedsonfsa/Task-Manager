package database

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func Connect(config DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%ss port=%d  user=%s password=%s dbname=%s sslmod=disable", config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}