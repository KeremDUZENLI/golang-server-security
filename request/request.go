package request

import (
	"fmt"
	"net/http"
	"seguro/env"
	"sync"
)

var count int

func SendRequest(wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()

	req, err := http.NewRequest("GET", env.URL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	for {
		if count >= env.NUMREQUEST {
			break
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		resp.Body.Close()

		requestCounter()
	}
}

func requestCounter() {
	count++
	fmt.Printf("\n%v. Request\t", count)
}
