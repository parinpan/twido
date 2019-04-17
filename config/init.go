/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T09:57:22+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-17T11:48:02+07:00
 */

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	. "twido/dataprovider"
)

// apps config's bootstrap
var TwidoConfig, TwidoConfigErr = NewConfiguration(ConfigurationOption{
	Environment: "production",
	BasePath:    "/root//arts/go-projects/src/twido/config", // absolute path to config directory
})

func NewConfiguration(co ConfigurationOption) (*Configuration, error) {
	configuration := &Configuration{}
	choosenCfgPath := fmt.Sprintf("%s/%s.json", co.BasePath, co.Environment)

	configFilePath, _ := filepath.Abs(choosenCfgPath)
	configFile, err := os.Open(configFilePath)

	if nil != err {
		log.Println("Could not open app's config file." + err.Error())
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

func IsTwitterCredentialSet() bool {
	elementsToCheck := []string{
		TwidoConfig.TwitterApiKey.Consumer,
		TwidoConfig.TwitterApiKey.ConsumerSecret,
		TwidoConfig.TwitterApiKey.Access,
		TwidoConfig.TwitterApiKey.AccessSecret,
	}

	for _, keyStr := range elementsToCheck {
		if "" != keyStr {
			return false
		}
	}

	return true
}

func IsRebrandlyCredentialSet() bool {
	return "" != TwidoConfig.UrlShortener.Rebrandly.ApiKey
}
