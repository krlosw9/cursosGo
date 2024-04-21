package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/krlosw9/cursosGo/api-go/class-3/model"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := NewResponse(Error, "Método no permitido", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := NewResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := NewResponse(Error, "Hubo un problema al crear la persona", nil)
		responseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := NewResponse(Message, "Persona creada correctamente", nil)
	responseJson(w, http.StatusCreated, response)

}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := NewResponse(Error, "Método no permitido", nil)
		responseJson(w, http.StatusMethodNotAllowed, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := NewResponse(Error, "El id debe ser un numero entero positivo", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}
	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := NewResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}
	err = p.storage.Update(ID, &data)
	if err != nil {
		response := NewResponse(Error, "Hubo un problema al actualizar la persona", nil)
		responseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := NewResponse(Message, "Persona actualizada correctamente", nil)
	responseJson(w, http.StatusOK, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := NewResponse(Error, "Método no permitido", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response := NewResponse(Error, "Hubo un problema al obtener todas las personas", nil)
		responseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := NewResponse(Message, "Ok", data)
	responseJson(w, http.StatusOK, response)
}
