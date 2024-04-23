package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func loginCient(url, email, password string) {
	login := Login{
		Email:    email,
		Password: password,
	}

	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error en marshak de login: %v", err)
	}

	resp := httpClient(http.MethodPost, url, data)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error en lectura del body %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Se esperaba codigo 200, se obtuvo: %d, respuesta: %s", resp.StatusCode, string(body))
	}

	fmt.Println(string(body))
}
