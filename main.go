package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restapimongo/db"
	"restapimongo/routes"

	"github.com/joho/godotenv"
)

func main() {
	router := routes.HandleRoutes()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading from .env file!")
	}

	db.ConnectDatabase()

	PORT := os.Getenv("PORT")
	fmt.Println("server running at PORT :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, &router))
}
