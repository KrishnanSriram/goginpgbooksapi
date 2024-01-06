package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title" binding:"required,min=10,max=200"`
	Author string `json:"author" binding:"required,min=5,max=50"`
	ISBN   string `json:"isbn" binding:"required,min=3,max=50"`
}

type BookInput struct {
	Title  string `json:"title" binding:"required,min=10,max=200"`
	Author string `json:"author" binding:"required,min=5,max=50"`
	ISBN   string `json:"isbn" binding:"required,min=3,max=50"`
}

func Validate(input BookInput) {
	//var validate *validator.Validate
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		//return err.Error()
	}
}

func (bi BookInput) NewBookFromInput() Book {
	return Book{
		ID:     createNewUUID(),
		Title:  bi.Title,
		Author: bi.Author,
		ISBN:   bi.ISBN,
	}
}

func createNewUUID() string {
	return uuid.New().String()
}
