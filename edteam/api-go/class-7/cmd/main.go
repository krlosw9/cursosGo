package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/krlosw9/cursosGo/api-go/class-7/authorization"
	"github.com/krlosw9/cursosGo/api-go/class-7/handler"
	"github.com/krlosw9/cursosGo/api-go/class-7/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("No se pudo cargar los certificados: %v", err)
	}

	// Crear un nuevo archivo de logs
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error al abrir el archivo de logs: %v", err)
	}
	defer logFile.Close()

	// Configurar el logger para que escriba en el archivo
	mw := io.MultiWriter(os.Stdout, logFile)
	logger := log.New(mw, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Configurar y lanzar el servidor Echo
	if err := runServer(logger); err != nil {
		logger.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

func runServer(logger *log.Logger) error {
	// Configurar el servidor Echo
	e, store := configureServer(logger)

	// Rutas
	handler.RouteLogin(e, &store)
	handler.RoutePerson(e, &store)

	// Iniciar el servidor
	return e.Start(":8080")
}

func configureServer(logger *log.Logger) (*echo.Echo, storage.Memory) {

	store := storage.NewMemory()

	// Configurar el servidor Echo
	e := echo.New()

	e.Use(middleware.Recover())
	// Middleware para registrar todas las solicitudes
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: logger.Writer(),
	}))

	// Middleware para manejar excepciones no capturadas
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					errMsg := fmt.Sprintf("Panic: %v", r)
					logger.Println(errMsg)
					c.Error(fmt.Errorf(errMsg))
				}
			}()
			return next(c)
		}
	})

	return e, store
}
