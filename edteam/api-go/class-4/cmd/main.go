package main

import "github.com/krlosw9/cursosGo/api-go/class-4/funciones"

func execute(name string, f func(string)) {
	f(name)
}

func main() {
	name := "Comunidad EDTeam"
	execute(name, funciones.MiddlewareLog(funciones.Saludar))
	execute(name, funciones.MiddlewareLog(funciones.Despedirse))
}
