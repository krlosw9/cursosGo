package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/krlosw9/cursosGo/api-go/class-7/authorization"
	"github.com/labstack/echo/v4"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		log.Printf("Inica: Petición: %q, método: %q", r.URL.Path, r.Method)
		f(w, r)
		log.Printf("Finaliza la petición: %v", time.Since(timeStart))
	}
}

func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "no permitido"})
		}

		return f(c)
	}
}
