package router

import (
	"devbookAPI/src/router/routes"

	"github.com/gorilla/mux"
)

type Router *mux.Router

// Gen: Generate an new mux router
func Gen() *mux.Router {
	router := mux.NewRouter()

	routes.Config(router)

	return router
}
