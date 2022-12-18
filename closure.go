package main

import "fmt"

func main()  {
	name := "Zidan"
	counter := 0

	increment := func ()  {
		name = "Zidanesc"
		fmt.Println("Incementing")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}