package main

import (
	"devbook-api/src/app/router"
	"log"
	"net/http"
)

func main() {
	router := router.GenerateRouter()

	log.Println("Server Listening...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
