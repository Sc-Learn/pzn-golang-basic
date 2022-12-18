package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main()  {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")		
	case bool:
		fmt.Println("value", value, "is bool")		
	case int:
		fmt.Println("value", value, "is int")		
	default:
		fmt.Println("Unknown type")
	}
}