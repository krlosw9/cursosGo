package handler

import (
	"net/http"

	"github.com/krlosw9/cursosGo/api-go/class-7/authorization"
	"github.com/krlosw9/cursosGo/api-go/class-7/model"
	"github.com/labstack/echo/v4"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {

	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, "estructura no valida", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if !isLoginValid(&data) {
		resp := NewResponse(Error, "usuario o contraseña no validos", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := NewResponse(Error, "No se pudo generar el token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}
	resp := NewResponse(Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "krlosw9@gmail.com" && data.Password == "123456"
}
