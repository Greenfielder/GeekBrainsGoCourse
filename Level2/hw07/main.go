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
	rv := reflect.ValueOf(values)

	for i := 0; i < elementOfIn.Type().NumField(); i++ {
		tempName := elementOfIn.Type().Field(i).Name
		newName := elementOfIn.FieldByName(tempName)

		for _, e := range rv.MapKeys() {
			v := rv.MapIndex(e)
			switch t := v.Interface().(type) {

			case int:
				if i == 1 {
					newName.SetInt(int64(t))
				}
			case string:
				if i == 0 {
					newName.SetString(t)
				}
			case bool:
				if i == 2 {
					newName.SetBool(t)
				}
			default:
				fmt.Println("not found")
			}
		}
	}
	return nil
}

//Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и values map[string]interface{}
//(key - название поля структуры, которому нужно присвоить value этой мапы).
//Необходимо по значениям из мапы изменить входящую структуру in с помощью пакета reflect.
//Функция может возвращать только ошибку error. Написать к данной функции тесты (чем больше, тем лучше - зачтется в плюс).

func main() {

	newValues := map[string]interface{}{
		"Name":   "Rick",
		"Cost":   100,
		"IsReal": true,
	}
	var inner = In{"John", 50, false}
	//Не получается присвоить новые значения, хотя на момент присваивания корректно отбражаются и старые и новые значения.
	customFunc(inner, newValues)
	fmt.Println(inner)
}
