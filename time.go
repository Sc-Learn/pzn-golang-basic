package main

import (
	"fmt"
	"time"
)

func main()  {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2022, 12, 30, 1, 1, 1, 1, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Date())
	fmt.Println(utc.Year())
	fmt.Println(utc.Month())
	fmt.Println(utc.Day())
	fmt.Println(utc.Hour())
	fmt.Println(utc.Minute())
	fmt.Println(utc.Second())
	fmt.Println(utc.Nanosecond())

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2022-12-12")
	fmt.Println(parse)
}