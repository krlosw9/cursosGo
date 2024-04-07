package main

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
