package main

// Home work 8

import (
	"flag"
	"fmt"
)

// type Configuration struct {
// 	Port        *int
// 	DbUrl       *string
// 	JaegerUrl   *string
// 	SentryUrl   *string
// 	KafkaBroker *string
// 	SomeAppId   *string
// 	SomeAppKey  *string
// }

var Port = flag.Int("port", 8080, "port to start server")
var DbUrl = flag.String("dburl", "sql://user:password@greenfielder:443", "DataBase Url")
var JaegerUrl = flag.String("jaegerurl", "https://jaeger", "Jaeger Url")
var SentryUrl = flag.String("sentryurl", "https://sentry", "Sentry Url")
var KafkaBroker = flag.String("kafkabroker", "", "Broker for Kafka")
var SomeAppId = flag.String("someappid", "555", "Some AppId")
var SomeAppKey = flag.String("someappkey", "unlock", "Some AppKey")

func main() {
	flag.Parse()
	fmt.Println("Server configuration: ")
	fmt.Printf(" Port: %d\n DB_url: %s\n JaegerURL: %s\n SentryURL: %s\n KafkaBroker: %s\n SomeAppId: %s\n SomeAppKey: %s\n",
		*Port, *DbUrl, *JaegerUrl, *SentryUrl, *KafkaBroker, *SomeAppId, *SomeAppKey)

}
