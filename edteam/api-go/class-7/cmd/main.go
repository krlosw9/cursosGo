package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/krlosw9/cursosGo/api-go/class-7/authorization"
	"github.com/krlosw9/cursosGo/api-go/class-7/handler"
	"github.com/krlosw9/cursosGo/api-go/class-7/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("No se pudo cargar los certificados: %v", err)
	}

	store := storage.NewMemory()

	e := echo.New()

	handler.RouteLogin(e, &store)
	handler.RoutePerson(e, &store)

	log.Println("Servidor corriendo en http://127.0.0.1:8080/")
	e.Logger.Fatal(e.Start(":8080"))
}
