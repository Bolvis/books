package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type teapot struct {
	Teapot string `json:"teapot"`
}

func makeCoffee(c *gin.Context) {
	c.JSON(http.StatusTeapot, teapot{Teapot: "teapot"})
}
