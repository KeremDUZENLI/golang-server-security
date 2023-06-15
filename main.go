package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"seguro/env"
	"seguro/request"
	"seguro/router"
)

var (
	wg            sync.WaitGroup
	client        = &http.Client{}
	storeChannels = make(chan int)
)

func main() {
	for i := 1; i <= env.NUMREQUEST; i++ {
		go requestSender()
		requestCounterToScreen(i)
	}

	env.PrintScan("PRESS ENTER TO START", nil)
	close(storeChannels)
	wg.Wait()
}

func init() {
	load()
	if env.URL == "" {
		env.LoadLocalEnvFile()
		go setupLocalServer()
	}
}

func load() {
	env.PrintScan("WELCOME TO THE HELL")

	env.PrintScan("URL (For Local Empty)", &env.URL)
	env.PrintScan("CONCURRENCY", &env.CONCURRENCY)
	env.PrintScan("NUMREQUEST", &env.NUMREQUEST)
}

func setupLocalServer() {
	err := router.SetupRouter().Run(env.PORT)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func requestSender() {
	wg.Add(env.CONCURRENCY)
	<-storeChannels

	defer wg.Done()
	request.SendRequest(&wg, client)
}

func requestCounterToScreen(i int) {
	time.Sleep(time.Second / 1000)
	fmt.Printf("\rThreads [%v] are ready", int(i))
	os.Stdout.Sync()
}
