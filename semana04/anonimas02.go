package main

import "fmt"

func main() {
	x := 0

	incremento := func() int {
		x++
		return x
	}
	incremento()
	incremento()
	fmt.Println(x)

}