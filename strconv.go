package main

import (
	"fmt"
	"strconv"
)

func main()  {
	boolean, err := strconv.ParseBool("true")
	showResult(boolean, err)
	boolean2, err := strconv.ParseBool("true mint")
	showResult(boolean2, err)

	number, err := strconv.ParseInt("1000000", 10, 32)
	showResult(number, err)
	number2, err := strconv.ParseInt("1000000a", 10, 32)
	showResult(number2, err)

	value := strconv.FormatInt(100000, 10)
	fmt.Println(value)

	valueInt, err := strconv.Atoi("2000000")
	showResult(valueInt, err)
}

func showResult(result interface{}, err error)  {
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}