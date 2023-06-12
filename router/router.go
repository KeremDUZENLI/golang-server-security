package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", handleRoot)
	return router
}

func handleRoot(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World")
}
