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

var portInit *int
var dbUrlInit *string
var someAppIdInit *string
var someAppKeyInit *string

func init() {

	portInit = flag.Int("port", 8080, "port to start server")
	dbUrlInit = flag.String("dburl", "sql://user:password@greenfielder:443", "DataBase Url")
	someAppIdInit = flag.String("someappid", "555", "Some AppId")
	someAppKeyInit = flag.String("someappkey", "unlock", "Some AppKey")
}

func LoadConfig() (*Configuration, error) {

	configuration := &Configuration{}

	port := portInit
	configuration.Port = int(*port)

	dbUrl := dbUrlInit
	configuration.DbUrl = *dbUrl

	someAppId := someAppIdInit
	configuration.SomeAppId = *someAppId

	someAppKey := someAppKeyInit
	configuration.SomeAppKey = *someAppKey

	return configuration, nil
}
