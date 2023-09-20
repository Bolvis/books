package api

import (
	"books/model"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	db := model.SetupModels()

	router.Use(func(c *gin.Context) {
		c.Set(dbKey, db)
		c.Next()
	})
	setupRoutes(router)

	if err := router.Run(strings.Join([]string{URL, port}, ":")); err != nil {
		fmt.Printf(err.Error(), gin.Logger())
	}
}
