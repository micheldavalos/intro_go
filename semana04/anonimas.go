package main

import "fmt"

func main() {
	sumar := func(x, y int) int {
		return x + y
	}

	fmt.Println(sumar(8, 9))
}