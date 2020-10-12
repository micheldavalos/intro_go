package main

import "fmt"

func promedio(arreglo []int) float64 { // se crea una copia
	total := 0.0
	for _, v := range arreglo {
		total += float64(v)
	}

	return total / float64(len(arreglo))
}

func main() {
	a := []int{1, 4, 2}
	fmt.Println(promedio(a))
}