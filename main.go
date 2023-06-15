package main

import (
	"fmt"
	"net/http"
	"sync"

	"seguro/env"
	"seguro/request"
	"seguro/router"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < env.NUMREQUEST; i++ {
		go requestSender()
	}

	wg.Wait()
}

func init() {
	env.Load()
	if env.DECIDE == 1 {
		setup()
	}
}

func setup() {
	r := router.SetupRouter()

	go func() {
		err := r.Run(env.PORT)
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()
}

func requestSender() {
	wg.Add(env.NUMREQUEST)

	client := &http.Client{}

	defer wg.Done()
	request.SendRequest(&wg, client)
}
