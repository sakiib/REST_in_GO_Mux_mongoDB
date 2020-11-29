package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restapimongo/db"
	"restapimongo/model"
	"restapimongo/routes"

	"github.com/joho/godotenv"
)

func main() {
	db.Books = []model.Book{
		model.Book{ID: "1", Isbn: "123", Title: "Book 1", Author: &model.Author{Firstname: "sakib", Lastname: "alamin"}},
		model.Book{ID: "2", Isbn: "234", Title: "Book 2", Author: &model.Author{Firstname: "nova", Lastname: "khan"}},
		model.Book{ID: "3", Isbn: "345", Title: "Book 3", Author: &model.Author{Firstname: "lionel", Lastname: "messi"}},
		model.Book{ID: "4", Isbn: "456", Title: "Book 4", Author: &model.Author{Firstname: "diego", Lastname: "maradona"}},
	}
	fmt.Println(db.Books)
	router := routes.HandleRoutes()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading from .env file!")
	}

	PORT := os.Getenv("PORT")
	fmt.Println("server running at PORT :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, &router))
}
