package helper

import "fmt"

var version = 1
var Application = "Learn golang"

func SayHello(name string)  {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string)  {
	fmt.Println("Goodbye", name)
}