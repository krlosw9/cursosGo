package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

}

func httpClient(method, url string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}

	return response
}
