package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Zidane"
	lastName = "Sc"

	return
}

func main()  {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}