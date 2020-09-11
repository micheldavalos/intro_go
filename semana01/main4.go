package main

import "fmt"

var nombre string = "Michel"
var global int64

func main() {
	fmt.Println(nombre)
	fmt.Println(global)
	const apellido string = "Davalos"
	fmt.Println(apellido)
	// apellido = "Boites"
}
