package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"seguro/common"
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
	for i := 1; i <= env.NUMBERREQUEST; i++ {
		go requestSender()
		requestCounterToScreen(i)
	}

	common.PrintScan("PRESS ENTER TO START", nil)
	close(storeChannels)
	wg.Wait()
}

func init() {
	env.LoadValuesGiven()
	if env.URL == "" {
		env.LoadValuesEnvFile()
		go runRouter()
	}
}

func runRouter() {
	err := router.SetupRouter().Run(env.PORT)
	common.PrintError(err)
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
