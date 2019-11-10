package database

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"golang-kafka-v5-crud/cmd/producer/config.go"
)

func ConnectPostgres() (*sql.DB, error) {
	conn, err := sql.Open(config.POSTGRESDRIVER, config.POSTGRESURL)
	if err != nil {
		logrus.Errorf("Cant connect to database got error: %v", err.Error())
		return nil, err
	}

	return conn, nil
}
