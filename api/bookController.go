package api

import (
	"books/model"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addBook(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var input model.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlInsert := `INSERT INTO "books" ("isbn", "title", "author") VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlInsert, input.ISBN, input.Title, input.Author)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, input)
}

func getAllBooks(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	rows, err := db.Query(`SELECT * FROM books`)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	defer rows.Close()
	var books []model.Book
	var isbn string
	var title string
	var author string
	var errors []string
	for rows.Next() {
		if err := rows.Scan(&isbn, &title, &author); err != nil {
			errors = append(errors, err.Error())
		}
		book := model.Book{ISBN: isbn, Title: title, Author: author}
		books = append(books, book)
	}

	if len(errors) != 0 {
		c.JSON(http.StatusInternalServerError, gin.H(map[string]any{"errors": errors}))
	}
	c.IndentedJSON(http.StatusOK, books)
}
