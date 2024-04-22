package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/krlosw9/cursosGo/api-go/class-3/authorization"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		log.Printf("Inica: Petición: %q, método: %q", r.URL.Path, r.Method)
		f(w, r)
		log.Printf("Finaliza la petición: %v", time.Since(timeStart))
	}
}

func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			forbidden(w, r)
			return
		}
		f(w, r)
	}
}

// Handler de sin autorizacion
func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No tiene autorización"))
}
