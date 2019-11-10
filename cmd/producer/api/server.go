package api

import (
	"golang-kafka-v5-crud/cmd/producer/api/routes"

	"golang-kafka-v5-crud/cmd/producer/config.go"
)

func Run() {
	router := routes.New()
	router.Run(":" + config.PORT)
}
