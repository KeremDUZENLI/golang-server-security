package router

import (
	"net/http"
	"strconv"
	"strings"

	"seguro/env"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", handleRoot)
	return router
}

func handleRoot(c *gin.Context) {
	env.COUNTER++
	env.LISTE = append(env.LISTE, env.COUNTER)

	numbers := make([]string, len(env.LISTE))
	for i, num := range env.LISTE {
		numbers[i] = strconv.Itoa(num)
	}

	response := strings.Join(numbers, ".Request\n")
	c.String(http.StatusOK, response)
}
