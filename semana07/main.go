package main

import (
	"fmt"
	"time"
)

// goroutine II
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * 100)
	}
}
//go routine I
func main() {
	go f(0)
	go f(1)
	go f(3)
	go f(4)

	var input string
	fmt.Scanln(&input)
}