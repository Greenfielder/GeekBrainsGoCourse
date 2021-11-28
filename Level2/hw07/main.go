package main

import (
	"fmt"
	"reflect"
)

type In struct {
	Name   string
	Cost   int
	IsReal bool
}

func customFunc(in In, values map[string]interface{}) error {
	valuesOfStrust := reflect.ValueOf(&in)
	elementOfIn := valuesOfStrust.Elem()

	for i := 0; i < elementOfIn.Type().NumField(); i++ {
		tempName := elementOfIn.Type().Field(i).Name
		newName := elementOfIn.FieldByName(tempName)
		fmt.Println(tempName, newName, "-----")
	}

	return nil
}

func main() {
	newValues := map[string]interface{}{
		"Name":   "Rick",
		"Cost":   100,
		"IsReal": true,
	}

	var inner In
	//На данный момент, я не понимаю в какую сторону дальше двигаться, что бы получить нужный результат
	customFunc(inner, newValues)
	fmt.Println(inner)
}
