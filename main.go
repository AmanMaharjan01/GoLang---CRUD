package main

import (
	"fmt"
	"go-postgres/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	r := router.Router()

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	fmt.Printf("Starting server on the %s.", port)

	listenerPort := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(listenerPort, r))
}
