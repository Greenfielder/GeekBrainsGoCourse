package config

import (
	"flag"
)

type Configuration struct {
	Port       int
	DbUrl      string
	SomeAppId  string
	SomeAppKey string
}

func LoadConfig() (*Configuration, error) {

	// var Port = flag.Int("port", 8080, "port to start server")
	// var DbUrl = flag.String("dburl", "sql://user:password@greenfielder:443", "DataBase Url")
	// var SomeAppId = flag.String("someappid", "555", "Some AppId")
	// var SomeAppKey = flag.String("someappkey", "unlock", "Some AppKey")
	// fmt.Println("Server configuration: ")
	// fmt.Printf(" Port: %d\n DB_url: %s\n  SomeAppId: %s\n SomeAppKey: %s\n", *Port, *DbUrl, *SomeAppId, *SomeAppKey)
	configuration := &Configuration{}

	port := flag.Int("port", 8080, "port to start server")
	configuration.Port = int(*port)

	dbUrl := flag.String("dburl", "sql://user:password@greenfielder:443", "DataBase Url")
	configuration.DbUrl = *dbUrl

	someAppId := flag.String("someappid", "555", "Some AppId")
	configuration.SomeAppId = *someAppId

	someAppKey := flag.String("someappkey", "unlock", "Some AppKey")
	configuration.SomeAppKey = *someAppKey

	return configuration, nil
}
