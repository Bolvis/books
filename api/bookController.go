package api

import (
	"books/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func addBook(c *gin.Context) {
	db := c.MustGet(dbKey).(*gorm.DB)

	var input model.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isAuthorValid(input.Author) {
		c.JSON(http.StatusBadRequest, "The authors name has to start with letter A")
		return
	}

	if err := db.Create(input).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, input)
}

func getAllBooks(c *gin.Context) {
	db := c.MustGet(dbKey).(*gorm.DB)

	var books []model.Book
	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}

func isAuthorValid(author string) bool {
	return strings.HasPrefix(strings.ToLower(author), "a")
}
