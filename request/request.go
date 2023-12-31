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

	response, err := httpResponse(client, request)
	common.PrintError(err)

	err = httpResponseBody(response)
	common.PrintError(err)
}

func httpRequest() (*http.Request, error) {
	request, err := http.NewRequest("GET", env.URL, nil)
	if err != nil {
		return nil, errors.New("\n!failed to send request")
	}

	return request, nil
}

func httpResponse(client *http.Client, request *http.Request) (*http.Response, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("\n!failed to get response")
	}

	return response, nil
}

func httpResponseBody(response *http.Response) error {
	if responseBody := response.Body.Close(); responseBody != nil {
		return errors.New("\n!failed to get responseBody")
	}

	return nil
}
