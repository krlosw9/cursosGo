package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 2)

	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c))

	close(c)

	for v := range c {
		fmt.Println(v)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Mensaje1 ", email1)
	go message("Mensaje2 ", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email ", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email ", m2)
		}
	}

	close(email1)
	close(email2)

}
