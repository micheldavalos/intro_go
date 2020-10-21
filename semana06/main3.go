package main

import (
	"fmt"
	"os"
)

func main()  {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	total := stat.Size()
	bs := make([]byte, total)
	count, err := file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(bs)
	fmt.Println(str, "bytes: ", count)

}