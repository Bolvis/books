package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func setupRoutes(e *gin.Engine) {
	e.GET(strings.Join([]string{basePath, "getAllBooks"}, "/"), getAllBooks)

	e.POST(strings.Join([]string{basePath, "addBook"}, "/"), addBook)
}
