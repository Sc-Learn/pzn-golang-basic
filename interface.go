package main

import "fmt"

type HasName interface{
	getName() string
}

func SayHello(hasName HasName)  {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main()  {
	var zidan Person
	zidan.Name = "Zidan"
	SayHello(zidan)

	var cat Person
	cat.Name = "puss"
	SayHello(cat)
}