package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Aldi")
	data.PushBack("Syah")
	data.PushBack("Putra")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // aldi

	next := head.Next() // syah
	fmt.Println(next.Value)

	next = next.Next() // putra
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
