package routes

import (
	"devbookAPI/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route: Represents an route from API
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	Auth    bool
}

func Config(router *mux.Router) {
	// Popula variavel app_routescomtodas as rotas do app
	origins := [][]Route{userRoutes, authRoutes, postRoutes}

	var app_routes []Route

	func() {
		for _, v := range origins {
			app_routes = append(app_routes, v...)
		}
	}()

	// start routes
	for _, route := range app_routes {
		if route.Auth {
			router.HandleFunc(route.URI, middlewares.Auth(route.Handler)).Methods(route.Method)

		} else {
			router.HandleFunc(route.URI, route.Handler).Methods(route.Method)

		}

	}
}
