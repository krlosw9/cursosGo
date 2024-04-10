package main

import "fmt"

func main() {
	tasks := []int{2, 34, 36, 17, 10, 40, 6, 22, 41, 44, 33, 2, 5, 2, 34, 36, 17, 10, 40, 6, 22, 41, 44, 33, 2, 5, 17, 10, 40, 6, 22, 2, 5, 2, 34, 36, 22, 41, 44}
	nWorkers := 5
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started, job with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d finished, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
