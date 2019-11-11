package v1

import (
	"golang-kafka-v5-crud/cmd/producer/api/controllers"
	"net/http"
)

var packageUser = "user"
var UserRouter = []Route{
	Route{
		Uri:          "/list",
		Method:       http.MethodGet,
		Handler:      controllers.ListUser,
		AuthRequired: false,
		Package:      packageUser,
	},
	Route{
		Uri:          "/detail/:id",
		Method:       http.MethodGet,
		Handler:      controllers.DetailUser,
		AuthRequired: false,
		Package:      packageUser,
	},
}
