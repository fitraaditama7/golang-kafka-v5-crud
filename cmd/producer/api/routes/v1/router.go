package v1

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Uri          string
	Method       string
	Handler      func(*gin.Context)
	AuthRequired bool
	Package      string
}

func Load() []Route {
	var routes []Route
	routes = append(routes, UserRouter...)
	return routes
}
