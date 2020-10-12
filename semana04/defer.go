package main

import "fmt"

func primero() {
	fmt.Println("primero")
}

func segundo() {
	fmt.Println("segundo")
}

func main() {
	defer primero()
	segundo()
}