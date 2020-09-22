package main

import "fmt"

func main()  {
	var temp int
	fmt.Print("Temp: ")
	fmt.Scan(&temp)
	if temp < 0 {
		fmt.Println("Está helado")
	} else if temp >= 0 && temp < 12 {
		fmt.Println("Está haciendo frío")
	} else if temp >= 12 && temp < 18 {
		fmt.Println("Está agusto")
	} else {
		fmt.Println("Está calurso")
	}

}