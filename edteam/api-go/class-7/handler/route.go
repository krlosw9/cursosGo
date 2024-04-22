package handler

import (
	"github.com/krlosw9/cursosGo/api-go/class-7/middleware"
	"github.com/labstack/echo/v4"
)

func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	person := e.Group("/v1/persons")
	person.Use(middleware.Authentication)

	person.POST("", h.create)
	person.PUT("/:id", h.update)
	person.GET("", h.getAll)
	// mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	// mux.HandleFunc("/v1/persons/get", h.getByID)
}

func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v1/login", h.login)
}
