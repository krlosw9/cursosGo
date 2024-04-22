package handler

import (
	"net/http"
	"strconv"

	"github.com/krlosw9/cursosGo/api-go/class-7/model"
	"github.com/labstack/echo/v4"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

// Create
func (p *person) create(c echo.Context) error {

	data := model.Person{}
	err := c.Bind(&data)
	if err != nil {
		response := NewResponse(Error, "La persona no tiene una estructura correcta", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := NewResponse(Error, "Hubo un problema al crear la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := NewResponse(Message, "Persona creada correctamente", nil)
	return c.JSON(http.StatusCreated, response)

}

// update
func (p *person) update(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := NewResponse(Error, "El id debe ser un numero entero positivo", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := model.Person{}
	err = c.Bind(&data)
	if err != nil {
		response := NewResponse(Error, "La persona no tiene una estructura correcta", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = p.storage.Update(ID, &data)
	if err != nil {
		response := NewResponse(Error, "Hubo un problema al actualizar la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := NewResponse(Message, "Persona actualizada correctamente", nil)
	return c.JSON(http.StatusOK, response)
}

// func (p *person) delete(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodDelete {
// 		response := NewResponse(Error, "Método no permitido", nil)
// 		responseJson(w, http.StatusMethodNotAllowed, response)
// 		return
// 	}

// 	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil {
// 		response := NewResponse(Error, "El id debe ser un numero entero positivo", nil)
// 		responseJson(w, http.StatusBadRequest, response)
// 		return
// 	}
// 	err = p.storage.Delete(ID)
// 	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
// 		response := NewResponse(Error, "El ID de la persona no existe", nil)
// 		responseJson(w, http.StatusBadRequest, response)
// 		return
// 	}
// 	if err != nil {
// 		response := NewResponse(Error, "Hubo un problema al eliminar el registro la persona", nil)
// 		responseJson(w, http.StatusInternalServerError, response)
// 		return
// 	}
// 	response := NewResponse(Message, "Registro eliminado correctamente", nil)
// 	responseJson(w, http.StatusOK, response)
// }

// func (p *person) getByID(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		response := NewResponse(Error, "Método no permitido", nil)
// 		responseJson(w, http.StatusMethodNotAllowed, response)
// 		return
// 	}

// 	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil {
// 		response := NewResponse(Error, "El id debe ser un numero entero positivo", nil)
// 		responseJson(w, http.StatusBadRequest, response)
// 		return
// 	}
// 	// person := person{}
// 	person, err := p.storage.GetByID(ID)
// 	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
// 		response := NewResponse(Error, "El ID de la persona no existe", nil)
// 		responseJson(w, http.StatusBadRequest, response)
// 		return
// 	}
// 	if err != nil {
// 		response := NewResponse(Error, "Hubo un problema consultar persona", nil)
// 		responseJson(w, http.StatusBadRequest, response)
// 		return
// 	}
// 	response := NewResponse(Message, "Ok", person)
// 	responseJson(w, http.StatusOK, response)
// }

// getAll
func (p *person) getAll(c echo.Context) error {

	data, err := p.storage.GetAll()
	if err != nil {
		response := NewResponse(Error, "Hubo un problema al obtener todas las personas", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := NewResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}
