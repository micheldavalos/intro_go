package main

import "fmt"

func incrementa(x *int) { 
	*x = *x + 1
}

func main() {
	x := 10
	incrementa(&x)
	fmt.Println(x)
}