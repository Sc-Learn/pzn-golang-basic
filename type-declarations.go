package main

import "fmt"

func main()  {
	type NoKtp string
	type Married = bool
	
	var noKtpZidan NoKtp = "1234312123431123"
	var marriedStatus Married = false
	fmt.Println(noKtpZidan)
	fmt.Println(marriedStatus)
}