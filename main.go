package main

import (
	"fmt"
	"net/http"
	"sync"

	"seguro/env"
	"seguro/request"
	"seguro/router"
)

var (
	wg     sync.WaitGroup
	decide int
)

func main() {
	for i := 0; i < env.NUMREQUEST; i++ {
		requestSender()
	}

	wg.Wait()
}

func init() {
	load()
}

func load() {
	env.PrintScan("WELCOME TO THE HELL")
	env.PrintScan("PRESS 1 FOR LOCAL TEST", &decide)

	if decide == 1 {
		env.LoadLocalEnvFile()
		setup()
	} else {
		env.PrintScan("URL", &env.URL)
	}

	env.PrintScan("CONCURRENCY", &env.CONCURRENCY)
	env.PrintScan("NUMREQUEST", &env.NUMREQUEST)
	env.PrintScan("PRESS ENTER", nil)
}

func setup() {
	err := router.SetupRouter().Run(env.PORT)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func requestSender() {
	wg.Add(env.NUMREQUEST)

	client := &http.Client{}

	defer wg.Done()
	request.SendRequest(&wg, client)
}
