package handlers

import (
	"booksapi/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type BookHandler struct {
	DB *sql.DB
}

func (bh BookHandler) GetBooks(c *gin.Context) {
	log.Println("GetBooks - bookhandler")
	rows, err := bh.DB.Query("SELECT id, title, author, isbn FROM books")
	if err != nil {
		log.Println("ERROR: Executing query \n%v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN); err != nil {
			log.Println("ERROR: Reading rows \n%v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning book",
				"message": err.Error()})
			return
		}
		books = append(books, b)
	}

	//if err := rows.Err(); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over books"})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"message": "List all books", "books": books})
}

func (bh BookHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	message := fmt.Sprintf("Find book %v", id)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (bh BookHandler) AddBook(c *gin.Context) {
	log.Println("Add a new book")
	// convert BookDTO into Book
	var newBookInput models.BookInput
	if err := c.ShouldBindJSON(&newBookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(newBookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newBook = newBookInput.NewBookFromInput()
	_, err := bh.DB.Exec("INSERT INTO books (id, title, author, isbn) VALUES ($1, $2, $3, $4)",
		newBook.ID, newBook.Title, newBook.Author, newBook.ISBN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the book",
			"errorMessage": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newBook)
}

func transformValidationErrors(err error) {

}
