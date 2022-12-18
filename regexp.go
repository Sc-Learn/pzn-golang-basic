package main

import (
	"fmt"
	"regexp"
)

func main()  {
	var regex regexp.Regexp = *regexp.MustCompile("z([a-z])n")

	fmt.Println(regex.MatchString("zan"))
	fmt.Println(regex.MatchString("zAn"))
	fmt.Println(regex.MatchString("zun"))

	search := regex.FindAllString("zan zAn zun", -1)
	fmt.Println(search)
}