package main

import "fmt"

func f3() (int, int) { // retornar varios valores
	return 5, 7
}

func f4() (a int, b int) {
	a = 0
	b = 10
	return a, b
}

func main() {
	fmt.Println(f3())
	fmt.Println(f4())

	a, b := f4()
	fmt.Println(a, b)
}