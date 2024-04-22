package handler

import (
	"encoding/json"
	"net/http"

	"github.com/krlosw9/cursosGo/api-go/class-3/authorization"
	"github.com/krlosw9/cursosGo/api-go/class-3/model"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := NewResponse(Error, "Método no permitido", nil)
		responseJson(w, http.StatusMethodNotAllowed, response)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := NewResponse(Error, "estructura no valida", nil)
		responseJson(w, http.StatusBadRequest, resp)
		return
	}
	if !isLoginValid(&data) {
		resp := NewResponse(Error, "usuario o contraseña no validos", nil)
		responseJson(w, http.StatusBadRequest, resp)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := NewResponse(Error, "No se pudo generar el token", nil)
		responseJson(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token": token}
	resp := NewResponse(Message, "Ok", dataToken)
	responseJson(w, http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "krlosw9@gmail.com" && data.Password == "123456"
}
