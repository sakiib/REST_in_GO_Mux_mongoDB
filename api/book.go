package api

import (
	"fmt"
	"net/http"
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
}

// UpdateBook ...
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateBook")
}

// DeleteBook ...
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteBook")
}
