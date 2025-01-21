package main

import (
	"fmt"
	"go-postgres/router"
	"log"
	"net/http"
	"os" // used to read the environment variable

	"github.com/joho/godotenv" // package used to read the .env file
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
