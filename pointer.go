package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address)  {
	address.Country = "awdwadwd"
}

func main()  {
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address2 := &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", ""}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	address := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	var addressPointer *Address = &address
	ChangeCountryToIndonesia(addressPointer)
	fmt.Println(address)
}