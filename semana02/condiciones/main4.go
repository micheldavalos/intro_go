package main

import "fmt"

func main()  {
	var temp int
	fmt.Print("Temp: ")
	fmt.Scan(&temp)

	switch {
	case temp < 0:
		fmt.Println("Está Helado")
	case temp >= 0 && temp < 12:
		fmt.Println("Está haciendo frío")
	case temp >= 12 && temp < 18:
		fmt.Println("Está agusto")
	default:
		fmt.Println("Está caluroso")
	}
}