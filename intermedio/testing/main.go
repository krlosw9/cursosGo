package main

import (
	"math"
)

func Sum(x, y int) int {
	return x + y
}

// go test -cover
// go test -coverprofile=coverage.out
// go tool cover -func=coverage.out
// go tool cover -html=coverage.out
func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// go test -cpuprofile=cpu.out
// go tool pprof cpu.out
// (pprof) top
// (pprof) list <nombre_funcion> -> list Fibonacci
// (pprof) web (si se obtiene el error: failed to execute dot. Is Graphviz installed? Error: exec: "dot": executable file not found in $PATH )
// ejecutar: sudo apt install graphviz
// (pprof) pdf
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func ComputeOnceFibonacci(n int) int {
	size := math.Max(2, float64(n+1))
	values := make([]int, int(size))
	values[0] = 0
	values[1] = 1

	for i := 2; i <= n; i++ {
		values[i] = values[i-1] + values[i-2]
		// fmt.Println(i, ":", values[i-1], "+", values[i-2])
	}

	return values[n]
}

// func main() {
// 	ComputeOnceFibonacci(50) // Probando la eficiencia de esta solucion de fibonacci
// }
