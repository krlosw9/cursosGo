package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "http://localhost:8080"

func main() {
	lc := loginClient(url+"/v1/login", "krlosw9@gmail.com", "123456")
	fmt.Println(lc)
}

func httpClient(method, url string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}

	return response
}
