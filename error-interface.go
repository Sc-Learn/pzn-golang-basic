package main

import (
	"errors"
	"fmt"
)

func Divide(number int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Divider can't 0")
	} else {
		result := number / divider
		return result, nil
	}
}

func main()  {
	result, err := Divide(100, 0)
	if err == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}