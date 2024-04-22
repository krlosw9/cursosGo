package main

import (
	"log"
	"net/http"

	"github.com/krlosw9/cursosGo/api-go/class-3/authorization"
	"github.com/krlosw9/cursosGo/api-go/class-3/handler"
	"github.com/krlosw9/cursosGo/api-go/class-3/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("No se pudo cargar los certificados: %v", err)
	}

	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	log.Println("Servidor corriendo en http://127.0.0.1:8080/")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
