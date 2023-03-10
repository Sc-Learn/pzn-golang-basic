package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type AnotherSample struct {
	Name string `required:"false" max:"10"`
	Description string `required:"false" max:"10"`
}

func IsValid(data any) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main()  {
	sample := Sample{"Zidan"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	sample.Name = ""
	fmt.Println(IsValid(sample))
	
	anotherSample := AnotherSample{"Zidan", ""}
	fmt.Println(IsValid(anotherSample))
}