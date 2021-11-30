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
	valuesOfStrust := reflect.ValueOf(&in).Elem()
	rv := reflect.ValueOf(values)

	fmt.Println(valuesOfStrust.CanSet())
	fmt.Println(valuesOfStrust.Field(0), valuesOfStrust.Field(1), valuesOfStrust.Field(2))

	for i := 0; i < valuesOfStrust.Type().NumField(); i++ {
		tempName := valuesOfStrust.Type().Field(i).Name
		newName := valuesOfStrust.FieldByName(tempName)

		for _, e := range rv.MapKeys() {
			v := rv.MapIndex(e)
			fmt.Print(v)
			switch t := v.Interface().(type) {
			case string:
				if i == 0 {
					fmt.Println(newName.CanSet())
					newName.SetString(t)
				}
			case int:
				if i == 1 {
					fmt.Println(newName.CanSet())
					newName.SetInt(int64(t))
				}
			case bool:
				if i == 2 {
					fmt.Println(newName.CanSet())
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
