package main

import (
	"fmt"
	"sync"
	"time"
)

// Buffered channels como semáforos
func main() {
	c := make(chan int, 2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomeThing(i, &wg, c)
	}
	wg.Wait()
}

func doSomeThing(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d finished\n", i)
	<-c
}
