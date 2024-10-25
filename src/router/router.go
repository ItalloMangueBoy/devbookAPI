package router

import (
	"devbookAPI/src/router/route"

	"github.com/gorilla/mux"
)

type Router *mux.Router

// Gen: Generate an new mux router
func Gen() *mux.Router {
	router := mux.NewRouter()

	route.Config(router)

	return router
}
