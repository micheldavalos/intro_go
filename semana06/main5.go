package main

import (
	"fmt"
	"sort"
)

func main()  {
	s := []int{8, 7, 6, 4}
	sort.Ints(s)
	fmt.Println(s)

	s1 := []int{1, 2, 3, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(s1)))
	fmt.Println(s1)
}