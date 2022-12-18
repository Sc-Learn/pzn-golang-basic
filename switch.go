package main

import "fmt"

func main()  {
	name := "Scc"

	switch name {
	case "Zidan":
		fmt.Println("Zidan")
	case "Sc":
		fmt.Println("Sc")
	default:
		fmt.Println("i dont know who you are")
	}

	switch length := len(name); length > 2 {
	case true:
		fmt.Println("Your name is too long")
	case false:
		fmt.Println("Hey you have a good name")
	}

	length := len(name)
	switch {
	case length < 2:
		fmt.Println("Your name is too short")
	case length > 2:
		fmt.Println("Your name is too long")
	default:
		fmt.Println("Hey you have a good name")
	}
}