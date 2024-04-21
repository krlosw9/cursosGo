package middleware

import (
	"log"
	"net/http"
	"time"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		log.Printf("Inica: Petición: %q, método: %q", r.URL.Path, r.Method)
		f(w, r)
		log.Printf("Finaliza la petición: %v", time.Since(timeStart))
	}
}
