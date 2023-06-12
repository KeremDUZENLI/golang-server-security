package request

import (
	"fmt"
	"net/http"
	"seguro/env"
	"sync"
)

func SendRequest(wg *sync.WaitGroup, client *http.Client, count *int) {
	defer wg.Done()

	req, err := http.NewRequest("GET", env.URL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	for {
		if *count >= env.NUMREQUEST {
			break
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		resp.Body.Close()

		requestCounter(count)
	}
}

func requestCounter(count *int) {
	*count++
	fmt.Printf("%v. Request\t", *count)
}
