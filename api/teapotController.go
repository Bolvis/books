package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func makeCoffee(c *gin.Context) {
	c.Status(http.StatusTeapot)
}
