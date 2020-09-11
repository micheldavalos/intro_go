package main

import "fmt"


func main() {
	var f float64

	fmt.Println("Float: ")
	// fmt.Scanf("%f", &f)
	fmt.Scan(&f)

	output := f * 2
	
	fmt.Println("Output: ", output)
}
