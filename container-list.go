package main

import (
	"container/list"
	"fmt"
)

func main()  {
	data := list.New()
	fmt.Println(*data)

	data.PushBack("Zidane")
	fmt.Println(*data)

	data.PushBack("Sc")
	fmt.Println(*data)

	data.PushFront("Muhammad")

	fmt.Println(*data)
	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for i := data.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}