package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 1)

	go say("Hello", c)
	fmt.Println(<-c)
	fmt.Println("World")
}
