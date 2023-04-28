package main

import (
	"log"
	"net/http"

	"github.com/ZiganshinDev/CRUD/router"
)

func main() {
	r := router.Router()

	log.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
