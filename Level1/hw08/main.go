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

// я сдаюсь... у меня не получается организовать пакеты.
// сначала перенёс все файлы в соответствии с  корневыми директориями.
// Теперь выдает "import cycle not allowed". Это при импорте в виде   _ "github.com/GreenFielder/GeekBrainsGoCourse/Level1/hw08",
// если без нижней черты  перед импортом, то сразу пропадает.
// Убилось просто два часа в пустую с этой организацией файлов.
// Сделал всё тапорно и без валидации. Если подскажете как победить ошибку с пакетом, то обязательно доделаю.

func main() {
	flag.Parse()
	fmt.Println("Server configuration: ")
	fmt.Printf(" Port: %d\n DB_url: %s\n JaegerURL: %s\n SentryURL: %s\n KafkaBroker: %s\n SomeAppId: %s\n SomeAppKey: %s\n",
		*Port, *DbUrl, *JaegerUrl, *SentryUrl, *KafkaBroker, *SomeAppId, *SomeAppKey)

}
