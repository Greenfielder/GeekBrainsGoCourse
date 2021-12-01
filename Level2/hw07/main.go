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

	inStruct := reflect.ValueOf(&in).Elem()

	for i := 0; i < inStruct.Type().NumField(); i++ {
		strKey := inStruct.Type().Field(i).Name
		strValue := inStruct.FieldByName(strKey)
		fmt.Println(strKey, strValue) // Здесь корректно выводит ключ и значение из структуры

		for x, y := range values {
			if x == strKey && strValue.Type().AssignableTo(reflect.TypeOf(y)) {
				fmt.Println(y, strValue.CanSet()) // Здесь корректно выводит значение мапы и факт того, что к полю структуры можно присвоить значение
				strValue.Set(reflect.ValueOf(y))
				fmt.Println(strValue, " - new value") // тут все присваивается
			}
		}
		fmt.Println(in) // Здесь видим, что структура полностью поменялась, но в мэйне мы видим, что ни чего не поменялось!!! :(((
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
