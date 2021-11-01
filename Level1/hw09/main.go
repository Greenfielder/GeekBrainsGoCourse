package main

// Home work 9

import (
	"flag"
	"fmt"
	"green/config"
	"log"
)

func main() {
	flag.Parse()

	configure, err := config.LoadConfig()
	if err != nil {
		log.Fatal("some info", err)
	}
	fmt.Println("Configuration file successfully load")
	fmt.Println("Server configuration: ")
	fmt.Printf(" Port: %d\n DB_url: %s\n SomeAppId: %s\n SomeAppKey: %s\n", configure.Port, configure.DbUrl,
		configure.SomeAppId, configure.SomeAppKey)

}
