package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

// Inicializar un módulo:
// go mod init github.com/username/module

//  Descargar una dependencia
// go get github.com/donvito/hellomod

// Limpiar dependencias sin utilizar
// go mod tidy

// Información de los módulos cacheados
// go mod download -json

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Platzi ")
}
