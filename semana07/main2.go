package main

import (
	"fmt"
	"time"
)

func pinger(c chan string)  {
	for {
		c <- "ping"
	}
}

func ponger(c chan string)  {
	for {
		c <- "pong"
	}
}

func printer(c chan string)  {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}