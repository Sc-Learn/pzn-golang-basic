package main

import (
	"fmt"
)

func main()  {
	counter := 1

	for counter <= 10 {
		fmt.Println("Counter: ", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Counter: ", counter)
	}

	slice := []string{"Muhammad", "Zidane", "Sc"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, name := range slice {
		fmt.Println("Index ", i, "Name: ", name)
	}

	person := make(map[string]string)
	person["firstName"] = "Muhammad"
	person["lastName"] = "Zidane Sc"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}