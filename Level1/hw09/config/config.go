package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
)

type Configuration struct {
	Port       int    `json:"port"`
	DbUrl      string `json:"dburl"`
	SomeAppId  string `json:"someappid"`
	SomeAppKey string `json:"someappkey"`
}

func validate(parsedUrl string) (*url.URL, error) {
	checkedUrl, err := url.Parse(parsedUrl)
	if err != nil || checkedUrl.Scheme == "" || checkedUrl.Host == "" {
		err = fmt.Errorf("%s is not valid url", parsedUrl)
		return nil, err
	}
	return checkedUrl, nil
}

func LoadConfig() (*Configuration, error) {

	var data Configuration
	var dbURL string
	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("Crush %s", err)
	}

	validate(data.DbUrl)
	if err != nil {
		log.Fatalf("Wrong url %s", err)
	} else {
		dbURL = data.DbUrl
	}

	configuration := &Configuration{}

	configuration.Port = int(data.Port)
	configuration.DbUrl = dbURL
	configuration.SomeAppId = data.SomeAppId
	configuration.SomeAppKey = data.SomeAppKey

	return configuration, nil
}
