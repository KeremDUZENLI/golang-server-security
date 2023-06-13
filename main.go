package main

import (
	"fmt"
	"net/http"
	"sync"

	"seguro/env"
	"seguro/request"
	"seguro/router"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(env.NUMREQUEST)

	client := &http.Client{}

	count := 0
	for i := 0; i < env.NUMREQUEST; i++ {
		go request.SendRequest(&wg, client, &count)
		if i%env.CONCURRENCY == 0 {
			wg.Wait()
		}
		fmt.Printf("\n%v. Request", i)
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
