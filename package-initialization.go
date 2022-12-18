package main

import (
	"fmt"
	"golang-pzn/basic-golang/database"
)

func main()  {
	database  := database.GetDatabase()
	fmt.Println(database)
}