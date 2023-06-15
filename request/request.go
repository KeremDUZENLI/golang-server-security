package request

import (
	"errors"
	"net/http"
	"seguro/common"
	"seguro/env"
	"sync"
)

func SendRequest(wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()

	request, err := httpRequest()
	common.PrintError(err)

	response := httpResponse(client, request)
	common.PrintError(response)
}

func httpRequest() (*http.Request, error) {
	request, err := http.NewRequest("GET", env.URL, nil)
	if err != nil {
		return nil, errors.New("\n!failed to send request")
	}

	return request, nil
}

func httpResponse(client *http.Client, request *http.Request) error {
	response, err := client.Do(request)
	if err != nil {
		return errors.New("\n!failed to get response")
	}

	if responseBody := response.Body.Close(); responseBody != nil {
		return errors.New("\n!failed to get responseBody")
	}

	return nil
}
