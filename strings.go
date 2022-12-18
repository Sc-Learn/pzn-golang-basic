package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("Zidane Sc", "Sc"))
	fmt.Println(strings.Contains("Zidane Sc", "Muhammad"))

	fmt.Println(strings.Split("Zidane Sc", " "))

	fmt.Println(strings.ToLower("Zidane SC "))
	fmt.Println(strings.ToUpper("Zidane SC"))
	
	fmt.Println(strings.Trim(" Zidane SC ", " "))

	fmt.Println(strings.ReplaceAll("Bonono ", "o", "a"))
}