package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/", saludar)
	e.GET("/dividir", dividir)

	handlerPersons := e.Group("/persons")
	handlerPersons.GET("/:id", getPersons)
	handlerPersons.POST("", create)
	handlerPersons.PUT("/:id", update)
	handlerPersons.DELETE("/:id", delete)

	e.Logger.Fatal(e.Start(":8080"))
}

func saludar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"saludo": "Hello, World!"})
}

func dividir(c echo.Context) error {
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)
	if f == 0 {
		return c.String(http.StatusBadRequest, "El valor no puede ser cero")
	}
	r := 3000 / f

	return c.String(http.StatusOK, strconv.Itoa(r))
}

func create(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{"message": "Creado"})
}

func getPersons(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"data": "Lista de personas..."})
}
func update(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{"message": "Actualizado id: " + id})
}
func delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{"message": "Eliminado id: " + id})
}
