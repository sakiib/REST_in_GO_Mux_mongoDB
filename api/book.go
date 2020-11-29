package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"restapimongo/db"
	"restapimongo/model"
	"time"
)

// GetBooks ...
func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBooks")
}

// GetBook ...
func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBook")
}

// CreateBook ...
func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateBook")
	w.Header().Add("Content-Type", "application/json")
	collection := db.Client.Database("goRESTapi").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	result, _ := collection.InsertOne(ctx, book)
	json.NewEncoder(w).Encode(result)
}

// UpdateBook ...
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateBook")
}

// DeleteBook ...
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteBook")
}
