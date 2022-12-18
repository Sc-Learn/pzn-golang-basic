package main

import "fmt"

func main()  {
	name1 := "zidan"
	name2 := "zidan"

	result := name1 == name2
	fmt.Println(result)

	value1 := 19
	value2 := 19

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}