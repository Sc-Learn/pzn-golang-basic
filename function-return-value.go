package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hey, what is your name"
	} else {
		return "Hello " + name
	}
}

func main()  {
	result := getHello("Zidanesc")
	fmt.Println(result)
	fmt.Println(getHello(""))
}