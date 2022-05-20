package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/silvergama/streetfair/app"
	"github.com/sirupsen/logrus"
)

var dbInstance *sql.DB

func SetupPostgreSQL() error {
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
	db, err := sql.Open("postgres", app.Config.Get("dbUrl"))
	if err != nil {
		logrus.Errorf("error when traying to connect database: %v", err)
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
