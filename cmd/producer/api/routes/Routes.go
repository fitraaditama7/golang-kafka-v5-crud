package routes

import (
	v1 "golang-kafka-v5-crud/cmd/producer/api/routes/v1"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	for _, route := range v1.Load() {
		version := r.Group(route.Package)

		switch route.Method {
		case "GET":
			version.GET(route.Uri, route.Handler)
		case "POST":
			version.GET(route.Uri, route.Handler)
		case "PUT":
			version.GET(route.Uri, route.Handler)
		case "PATCH":
			version.GET(route.Uri, route.Handler)
		case "DELETE":
			version.GET(route.Uri, route.Handler)
		case "OPTION":
			version.GET(route.Uri, route.Handler)
		}
	}
	return r
}
