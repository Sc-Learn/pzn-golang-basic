package main

import "fmt"

func main()  {
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Zidane"
	names[2] = "Sc"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		10,
		15,
		20,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))
}