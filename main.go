package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"seguro/env"
	"seguro/request"
	"seguro/router"
)

var wg sync.WaitGroup
var startChannel = make(chan int)

func main() {
	for i := 0; i < env.NUMREQUEST; i++ {
		time.Sleep(time.Microsecond * 50)
		wg.Add(1)
		go requestSender()
		fmt.Printf("\rThreads [%.0f] are ready", float64(i+1))
		os.Stdout.Sync()
	}

	fmt.Printf("\nPlease [Enter] for continue")
	_, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	close(startChannel)
	fmt.Println("lets begin...")
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
	<-startChannel
	client := &http.Client{}

	defer wg.Done()
	request.SendRequest(&wg, client)
}
