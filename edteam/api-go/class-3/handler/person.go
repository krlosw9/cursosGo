package handler

import (
	"encoding/json"
	"net/http"

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
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Método no permitido"}`))
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "La persona no tiene una estructura correcta"}`))
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "Hubo un problema al crear la persona"}`))
		return
	}

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type": "message", "message": "Persona creada correctamente"}`))

}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Método no permitido"}`))
		return
	}

	resp, err := p.storage.GetAll()
	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "Hubo un problema al obtener todas las personas"}`))
		return
	}

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "Hubo un problema al convertir el slice a json"}`))
		return
	}

}
