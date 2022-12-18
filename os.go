package main

import (
	"fmt"
	"os"
)

func main()  {
	args := os.Args
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(args[2])

	name, err := os.Hostname()
	if err == nil {
			fmt.Println(name)
		} else {
			fmt.Println("Error", err.Error())
	}

	npmToken := os.Getenv("NPM_TOKEN")
	fmt.Println("NPM_TOKEN", npmToken)
}