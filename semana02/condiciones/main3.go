package main

import "fmt"

func main()  {
	var op int
	
	fmt.Print("Opción: ")
	fmt.Scan(&op)

	switch op {
	case 1:
		fmt.Println("Opción 1")
	case 2:
		fmt.Println("Opción 2")
	case 3:
		fmt.Println("Opción 3")
	default:
		fmt.Println("Opción no existe")
	}
}