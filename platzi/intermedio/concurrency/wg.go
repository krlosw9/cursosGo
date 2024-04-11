package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomeThing(i, &wg)
	}

	wg.Wait()
}

func doSomeThing(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Stard %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("End %d\n", i)
}
