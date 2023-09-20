package api

import (
	"books/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	db := model.SetupModels()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	setupRoutes(router)

	router.Run(strings.Join([]string{"localhost", port}, ":"))
}
