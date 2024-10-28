package main

import (
	"devbookAPI/src/config"
	"devbookAPI/src/database"
	"devbookAPI/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	router := router.Gen()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), router))
}
