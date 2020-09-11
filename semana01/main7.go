package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	

	fmt.Print("Nombre: ")
	scanner.Scan()
	line = scanner.Text()

	fmt.Println("Hola", line)
}
