package main

import (
	"devbookAPI/src/router"
	"log"
	"net/http"
)

func main() {
	router := router.Gen()

	log.Fatal(http.ListenAndServe(":5000", router))
}
