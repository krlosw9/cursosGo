package handler

import (
	"net/http"

	"github.com/krlosw9/cursosGo/api-go/class-3/middleware"
)

func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", middleware.Log(middleware.Authentication(h.create)))
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/get-all", middleware.Log(h.getAll))
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/v1/persons/get", h.getByID)
}
