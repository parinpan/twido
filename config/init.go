/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T09:57:22+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-17T08:26:26+07:00
 */

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	. "twido/dataprovider"
)

func NewConfiguration(environment string) (*Configuration, error) {
	configuration := &Configuration{}
	configFilePath, _ := filepath.Abs("twido/config/" + environment + ".json")
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

// apps config's bootstrap
var TwidoConfig, TwidoConfigErr = NewConfiguration("production")
