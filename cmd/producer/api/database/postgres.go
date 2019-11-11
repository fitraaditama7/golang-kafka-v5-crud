package database

import (
	"database/sql"

	"golang-kafka-v5-crud/cmd/producer/config"

	"github.com/sirupsen/logrus"
	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	conn, err := sql.Open(config.POSTGRESDRIVER, config.POSTGRESURL)
	if err != nil {
		logrus.Errorf("Cant connect to database got error: %v", err.Error())
		return nil, err
	}

	return conn, nil
}
