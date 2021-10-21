package main

// Home work 8

import (
	"flag"
	"fmt"
	"green/config"
	"log"
)

// mickey@mickey-pc:~/go/src/github.com/GreenFielder/GeekBrainsGoCourse/Level1/hw08$ ./main
// 2021/10/21 22:42:29 &{8080 sql://user:password@greenfielder:443 555 unlock}
// mickey@mickey-pc:~/go/src/github.com/GreenFielder/GeekBrainsGoCourse/Level1/hw08$ ./main -port 2020
// flag provided but not defined: -port
// Usage of ./main:
// Разобрался с пакетами, то перерь не работает сама программа. Вернее как.
// Я разбил на пакеты, всё вынес в конфиг.
// Без флагов если запускать, то всё ок - выводит значения по умолчанию
// А если с флагами, то вылетает с сообщением "flag provided but not defined: -port"
// Я подозреваю что переменные надо вынести в init(), но тогда они видны только внутри инита.
// Нужна помощь!

func main() {
	flag.Parse()

	err, configure := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server configuration: ", configure)
}
