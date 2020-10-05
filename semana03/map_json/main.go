package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	materias := map[string]map[string]int{
		"Primero": {
			"Programación": 100,
		},
		"Tercero": {
			"Algoritmia": 90,
		},		
	}
	// _, err := materias["Estructura de Datos"]
	// if v, err := materias["Algoritmia"]; err  {
	// 	fmt.Println(v, err)
	// }
	

	j, _ := json.Marshal(materias)
	
	// fmt.Println(j)
	fmt.Println(string(j))
}