package main

import "fmt"

func main()  {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "May edited"
	// fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "New Month")
	fmt.Println(slice3)
	slice3[1] = "Not December"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Muhammad"
	newSlice[1] = "Zidane"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	checkCapSlice := make([]int, 2)
	checkCapSlice2 := []int{}
	fmt.Println(cap(checkCapSlice))
	fmt.Println(cap(checkCapSlice2))
}