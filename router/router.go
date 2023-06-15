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
	numbers := numbersListConvertString()
	response := strings.Join(numbers, ".Request\n")
	c.String(http.StatusOK, response)
}

func numbersListConvertString() []string {
	numbersList()

	numbers := make([]string, len(env.LISTE))
	for i, num := range env.LISTE {
		numbers[i] = strconv.Itoa(num)
	}
	return numbers
}

func numbersList() {
	env.COUNTER++

	if env.NUMREQUEST >= env.COUNTER {
		env.LISTE = append(env.LISTE, env.COUNTER)
		env.LOOPER = true
	}
}
