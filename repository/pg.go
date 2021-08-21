package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var dbInstance *sql.DB

func Setup() error {
	if dbInstance != nil {
		return nil
	}
	instance, err := setupDBInstance()
	if err != nil {
		return err
	}

	dbInstance = instance

	return nil
}

func setupDBInstance() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"unico",
		"123456",
		"localhost",
		"5432",
		"unico"))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GetInstance() *sql.DB {
	return dbInstance
}
