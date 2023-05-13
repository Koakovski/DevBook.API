package main

import (
	"devbook-api/src/app/router"
	"devbook-api/src/infra/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	router := router.GenerateRouter()

	stringApiPort := fmt.Sprintf(":%d", config.ApiPort)

	log.Println("Server Listening on", stringApiPort)
	log.Fatal(http.ListenAndServe(stringApiPort, router))
}
