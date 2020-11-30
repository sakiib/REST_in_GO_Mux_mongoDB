package db

import "restapimongo/model"

// Books ...
var Books = []model.Book{
	model.Book{Isbn: "123", Title: "Book 1", Author: &model.Author{Firstname: "sakib", Lastname: "alamin"}},
	model.Book{Isbn: "234", Title: "Book 2", Author: &model.Author{Firstname: "nova", Lastname: "khan"}},
	model.Book{Isbn: "345", Title: "Book 3", Author: &model.Author{Firstname: "lionel", Lastname: "messi"}},
	model.Book{Isbn: "456", Title: "Book 4", Author: &model.Author{Firstname: "diego", Lastname: "maradona"}},
}
