package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	. "twido/dataprovider"
)

func NewConfiguration(environment string) (*Configuration, error) {
	configuration := &Configuration{}
	configFile, err := os.Open("config/" + environment + "-config.json")

	if nil != err {
		log.Println("Could not open app's config file.")
		return nil, err
	}

	defer configFile.Close()
	configBytes, err := ioutil.ReadAll(configFile)

	if nil != err {
		log.Println("Could not read app's config file.")
		return nil, err
	}

	if err := json.Unmarshal(configBytes, &configuration); nil != err {
		log.Println("Could not load app's config file")
		return nil, err
	}

	return configuration, nil
}

// apps bootstrap
var TwidoConfig, TwidoConfigErr = NewConfiguration("production")
