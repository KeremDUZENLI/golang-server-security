package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	targetURL   = "http://localhost:8080/"
	numRequests = 1000
	concurrency = 100
)

var (
	port  = ":8080"
	count = 0
)

func main() {
	var wg sync.WaitGroup
	wg.Add(numRequests)

	client := &http.Client{}

	router := gin.Default()
	router.GET("/", handleRoot)

	go func() {
		err := router.Run(port)
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()

	for i := 0; i < numRequests; i++ {
		go sendRequest(&wg, client)
		if i%concurrency == 0 {
			wg.Wait()
		}
		fmt.Printf("%d. Request\n", i)
	}

	wg.Wait()
}

func sendRequest(wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	for {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		resp.Body.Close()
	}
}

func handleRoot(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
	requestCounter()
}

func requestCounter() {
	fmt.Printf("%v. Request\t", count)
	count++
}
