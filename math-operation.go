package main

import "fmt"

func main()  {
	result := 10 + 10
	fmt.Println(result)

	a := 10
	b := 10
	c := a * b
	fmt.Println(c)

	result += 80
	fmt.Println(result)

	result++
	fmt.Println(result)

	fmt.Println(-result)
	fmt.Println(+result)
}