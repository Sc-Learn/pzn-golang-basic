package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.4))

	fmt.Println(math.Floor(1.9))
	fmt.Println(math.Ceil(1.1))
	
	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}