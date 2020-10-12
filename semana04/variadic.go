package main

import "fmt"

// Variadic functions
func sumar(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}

	return total
}

func main() {
	a := []int{1, 4, 2}
	fmt.Println(sumar(4, 5, 2))
	fmt.Println(sumar(a...))
}