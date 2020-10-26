package main

import (
	"fmt"
	"strings"
)

func main()  {
	// fmt.Println(strings.Contains("distribuidos", "bui"))
	// fmt.Println(strings.Count("distribuidos", "i"))
	// fmt.Println(strings.HasPrefix("distribuidos", "di"))
	// fmt.Println(strings.HasSuffix("distribuidos", "os"))
	// fmt.Println(strings.Index("distribuidos", "as"))
	fmt.Println(strings.Join([]string{"Sistemas", "Distribuidos", "CUCEI"}, "\n"))
	// fmt.Println(strings.Repeat("distribudos", 2))
	// fmt.Println(strings.Replace("aaaaabbb", "a", "b", 2))
	// fmt.Println(strings.Split("Mi mama me mima", " "))
	// fmt.Println(strings.ToLower("Mi mama me mima"))
	// fmt.Println(strings.ToUpper("Mi mama me mima"))

	b := []byte("test")
	fmt.Println(b)

	s := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(s)
}