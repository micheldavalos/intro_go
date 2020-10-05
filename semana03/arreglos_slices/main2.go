package main

import "fmt"

func main() {
	var s []int
	fmt.Println(len(s), cap(s))
	s = append(s, 1, 4, 5, 6, 9)
	fmt.Println(len(s), cap(s))
	for _, v := range s {
		fmt.Println(v)
	}
}
