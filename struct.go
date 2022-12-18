package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string)  {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main()  {
	var zidan Customer
	zidan.Name = "Zidan"
	zidan.Address = "Wisma Asri"
	zidan.Age = 19

	// fmt.Println(zidan.Name)
	// fmt.Println(zidan.Address)
	// fmt.Println(zidan.Age)

	// wednesday := Customer{
	// 	Name: "Wednesday",
	// 	Address: "UK",
	// 	Age: 17,
	// }
	// fmt.Println(wednesday)

	// professor := Customer{"Profressor", "Spain/Korea", 40}
	// fmt.Println(professor)

	zidan.sayHello("Adams")
	zidan.sayHello("Bella")
}