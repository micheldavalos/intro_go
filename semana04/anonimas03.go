package main

import "fmt"

func generadorPares() func() uint {
	i := uint(0)
	return func() (par uint) {
		par = i
		i += 2
		return
	}
}

func main() {
	nextPar := generadorPares()
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())
}