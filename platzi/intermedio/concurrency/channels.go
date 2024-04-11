package main

import "fmt"

func concurrentRoutine(text string, c chan string) chan string {
	c <- text
	return c
}

// Canales Unbuffered: Estos canales no tienen capacidad de almacenamiento, lo que significa que solo pueden contener un valor a la vez. Cuando se envía un valor a través de un canal unbuffered, el proceso de envío se bloquea hasta que otro proceso recibe el valor del canal. Los canales unbuffered se utilizan para la sincronización de datos entre goroutines, asegurando que la comunicación se produzca de manera sincrónica.

// Canales Buffered: Estos canales tienen una capacidad definida y pueden contener varios valores antes de bloquear el envío. Cuando se envía un valor a través de un canal buffered, el envío no se bloquea a menos que el canal esté lleno. Esto permite que la comunicación entre goroutines sea asincrónica en cierta medida, ya que el envío no bloqueará la goroutine que lo realiza a menos que el canal esté lleno. Los canales buffered son útiles cuando la velocidad de producción y consumo de datos difiere entre goroutines, o cuando se necesita amortiguar la comunicación para reducir el bloqueo.

func main() {
	// ------------Channel unbuffered ---------------------
	c := make(chan string)                         // canal unbuffered | tamaño de cero
	go concurrentRoutine("Mensaje concurrente", c) // unbuffered bloquea esperando que alguien reciba el mensaje o la informacion, en este caso concurrentRoutine() es quien recibe el mensaje y desbloquea

	fmt.Println("Se recibe del canal unbuffered:", <-c)

	// ------------Channel Buffered ---------------------
	buf := make(chan int, 3) // Aqui el tamaño del canal es de 3 y este tipo de canal no bloquea

	buf <- 1
	buf <- 2
	buf <- 25 // Con este completa el cupo del canal que es 3

	fmt.Println("Canal Buffered 1:", <-buf)
	fmt.Println("Canal Buffered 2:", <-buf)
	fmt.Println("Canal Buffered 3:", <-buf)

}
