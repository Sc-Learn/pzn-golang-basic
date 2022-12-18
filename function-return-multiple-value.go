package main

import "fmt"

func getFullName() (string, string, int) {
	return "Muhammad", "Zidane Sc", 19
}

func main()  {
	firstName, lastName, _ := getFullName()
	fmt.Println(firstName, lastName)
}