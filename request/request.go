package request

import (
	"fmt"
	"net/http"
	"seguro/env"
	"sync"
)

func SendRequest(wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()

	req, err := http.NewRequest("GET", env.URL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	for env.LOOPER {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		resp.Body.Close()
	}
}
