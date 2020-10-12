package main

import "fmt"

func suma(x, y int) int { 
	return x + y
}

func main() {
	x := 10
	y := 1
	fmt.Println(suma(x, y))
}