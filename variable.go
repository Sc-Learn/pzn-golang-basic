package main

import "fmt"

func main() {
	var name string

	name = "Muhammad Zidane Sc"
	fmt.Println(name)

	name = "Zidane Sc"
	fmt.Println(name)

	// will be error
	// name = true
	// name = 12

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 19
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Muhammad"
		lastName = "Zidane Sc"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}