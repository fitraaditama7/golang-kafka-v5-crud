package controllers

import (
	"context"
	"golang-kafka-v5-crud/cmd/producer/api/database"
	"golang-kafka-v5-crud/cmd/producer/api/helper/response"
	"golang-kafka-v5-crud/cmd/producer/api/repositories"
	"golang-kafka-v5-crud/cmd/producer/api/repositories/repo"
	"net/http"
	"strconv"

	"golang-kafka-v5-crud/cmd/producer/config"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	var w = c.Writer
	var ctx = c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	db, err := database.ConnectPostgres()
	if err != nil {
		response.ErrorCustomStatus(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	PSQLRepository := repo.NewPSQLRepository(db)

	func(userPSQLRepository repositories.UserRepository) {
		ctx, cancel := context.WithTimeout(ctx, config.TIMEOUT)
		defer cancel()

		users, err := userPSQLRepository.List(ctx)
		if err != nil {
			response.ErrorCustomStatus(w, http.StatusInternalServerError, err)
			return
		}
		response.Response(w, http.StatusOK, "Success", users)
	}(PSQLRepository)
}

func DetailUser(c *gin.Context) {
	var w = c.Writer
	var ctx = c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorCustomStatus(w, http.StatusInternalServerError, err)
	}

	db, err := database.ConnectPostgres()
	if err != nil {
		response.ErrorCustomStatus(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	PSQLRepository := repo.NewPSQLRepository(db)

	func(userPSQPSQLRepository repositories.UserRepository) {
		ctx, cancel := context.WithTimeout(ctx, config.TIMEOUT)
		defer cancel()

		user, err := userPSQPSQLRepository.Detail(ctx, id)
		if err != nil {
			response.ErrorCustomStatus(w, http.StatusInternalServerError, err)
			return
		}
		response.Response(w, http.StatusOK, "Success", user)
	}(PSQLRepository)
}
