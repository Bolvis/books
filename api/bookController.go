package api

import (
	"books/model"
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addBook(c *gin.Context) {
	db := c.MustGet(dbKey).(*gorm.DB)

	var input model.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(input)

	c.IndentedJSON(http.StatusOK, input)
}

func getAllBooks(c *gin.Context) {
	db := c.MustGet(dbKey).(*gorm.DB)

	var books []model.Book
	db.Find(&books)

	c.IndentedJSON(http.StatusOK, books)
}
