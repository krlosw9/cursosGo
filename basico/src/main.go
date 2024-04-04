package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) { // Gorutine

	defer wg.Done() // Esta linea se va a ejecutar hasta el final de la funcion, y de esta forma libera el gorutine del WaitGroup
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup // El paquete sync permite interacturar de forma primitiva con las gorutine. Variable que acomula un conjunto de gorutines y los va liberando poco a poco

	wg.Add(1)            // Indicamos que vamos a agregar 1 Gorutine al WaitGroup para que espere su ejecucion antes de que la gurutine base (main) muera, y así le de tiempo a la siguiente gorutine de ejecutarse
	go say("Hello", &wg) // la palabra reservada go ejecutará la funcion de forma concurrente
	wg.Wait()            // Funcion del WaitGroup que sirve para decirle al gorutine principal (main) que espere hasta que todas las gorutine del WaitGroup finalicen, es decir, hasta que se ejecute 'defer wg.Done()' en cada una de las goroutines

	fmt.Println("world")

	go func(text string) { // Funciona anonima
		fmt.Println(text)
	}("adios")

	time.Sleep(time.Second * 1) // ! Funcion para que cuando llegue a esta linea espere el tiempo indicado (lo suficiente para que la Gorutine ejecute su funcion de forma concurrente)

	// time.Sleep(time.Second * 1)
	// Nota:
	// - Para fines educativos se hace uso de la funcion Sleep(), pero en realidad NO es una buena practica, es mejor utilizar los WaitGroups
	// - los WaitGroups son mas optimos pero son complicados de mantener por lo tanto se utiliza mas los channels

}
