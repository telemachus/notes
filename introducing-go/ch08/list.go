package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l list.List
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
