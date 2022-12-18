package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married()  {
	man.Name = "Mr. " + man.Name
}

func main()  {
	zidan := Man{Name: "Zidan"}
	zidan.Married()

	fmt.Println(zidan.Name)
}