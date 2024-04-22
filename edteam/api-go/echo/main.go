package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", saludar)
	e.Logger.Fatal(e.Start(":8080"))
}

func saludar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"saludo": "Hello, World!"})
}
