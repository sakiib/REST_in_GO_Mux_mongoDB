package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"restapimongo/db"
	"restapimongo/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// GetBooks ...
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var books []model.Book
	collection := db.Client.Database("goRESTapi").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message: "` + err.Error() + `"}"`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book model.Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message: "` + err.Error() + `"}"`))
		return
	}
	json.NewEncoder(w).Encode(books)
}

// GetBook ...
func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBook")
}

// CreateBook ...
func CreateBook(w http.ResponseWriter, r *http.Request) {
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
