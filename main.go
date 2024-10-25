package main

import (
	"devbookAPI/src/config"
	"devbookAPI/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	router := router.Gen()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), router))
}
