package main

import (
	"fmt"
	"os"
)

func main()  {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("Hola mundo")
}