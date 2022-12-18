package main

import (
	"flag"
	"fmt"
)

func main()  {
	host := flag.String("host", "localhost", "Put your database localhost")
	user := flag.String("user", "root", "Put your database user")
	password := flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "put your number")

	flag.Parse()

	fmt.Println("Host", *host)
	fmt.Println("user", *user)
	fmt.Println("password", *password)
	fmt.Println("number", *number)
}