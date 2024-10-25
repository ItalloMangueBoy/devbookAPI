package route

import (
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
	var routes = append([]Route{}, userRoutes...)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
}
