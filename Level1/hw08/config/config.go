package config

// Home work 8

type Configuration struct {
	Port        int
	DbUrl       string
	JaegerUrl   string
	SentryUrl   string
	KafkaBroker string
	SomeAppId   string
	SomeAppKey  string
}
