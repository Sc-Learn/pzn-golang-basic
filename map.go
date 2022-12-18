package main

import "fmt"

func main()  {
	person := map[string]string{
		"name": "zidan",
		"city": "Bekasi",
	}

	person["role"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["city"])
	fmt.Println(person["role"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Learn Golang for baby"
	book["author"] = "Zidane Sc"
	book["ups"] = "my fault"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}