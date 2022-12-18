package main

import "fmt"

func main()  {
	var name = "Sc"

	if name == "zidan" {
		fmt.Println("Hello, zidan")
	} else if name == "Sc" {
			fmt.Println("Hello, Sc")
	} else {
		fmt.Println("Hi, who are you?")
	}

	if length := len(name); length > 2 {
		fmt.Println("Your name is too long")
		} else {
		fmt.Println("the length of your name is enough")
	}
}