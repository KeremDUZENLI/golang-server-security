package router

import (
	"net/http"
	"strconv"
	"strings"

	"seguro/env"

	"github.com/gin-gonic/gin"
)

var (
	count int = 1
	liste []int
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

	numbers := make([]string, len(liste))
	for i, num := range liste {
		numbers[i] = strconv.Itoa(num)
	}

	return numbers
}

func numbersList() {
	if count <= env.NUMBERREQUEST {
		liste = append(liste, count)
		count++
	}
}
