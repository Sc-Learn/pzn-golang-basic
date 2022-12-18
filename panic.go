package main

import "fmt"

func endApp()  {
	message := recover()
	
	if message != nil {
		fmt.Println("Error with message: ", message)
	}

	fmt.Println("Application finished")
}

func runApplication(error bool)  {
	defer endApp()
	if error {
		panic("APPLICATION_HAVE_AN_ISSUE")
	}

	fmt.Println("Application work")
}

func main()  {
	runApplication(true)
	fmt.Println("You must see this message")
}