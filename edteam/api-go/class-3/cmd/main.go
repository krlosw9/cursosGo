package main

import (
	"log"
	"net/http"

	"github.com/krlosw9/cursosGo/api-go/class-3/handler"
	"github.com/krlosw9/cursosGo/api-go/class-3/storage"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	log.Println("Servidor corriendo en http://127.0.0.1:8080/")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
