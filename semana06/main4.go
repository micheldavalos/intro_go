package main

import (
	"fmt"
	"container/list"
)

func main()  {
	var l list.List

	l.PushBack(1)
	l.PushBack("string")
	l.PushFront(3.5)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println(l)

}