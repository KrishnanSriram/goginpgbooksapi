package main

import (
	"booksapi/handlers"
	"booksapi/intialize"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var dbstore intialize.PostgreDB

func init() {
	fmt.Println("Initialize appliction level ops")
	intialize.InitApp()
	var err error
	dbstore, err = intialize.InitDB()
	if err != nil {
		log.Fatal("Failed to open DB connection!! Terminate application execution")
	}

}

func main() {
	fmt.Println("Welcome to Books API")
	r := gin.Default()
	bh := handlers.BookHandler{DB: dbstore.DB}
	r.GET("/books", bh.GetBooks)
	r.GET("/books/:id", bh.GetBookByID)
	r.POST("/books", bh.AddBook)
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Println("ERROR: Cannot close DB connection. Possible MEMORY leak??")
		}
		log.Println("DB connection closed!!")
	}(dbstore.DB)
	r.Run()
}
