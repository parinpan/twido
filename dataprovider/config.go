package dataprovider

import (
	"net/url"
)

type TwitterAE TwitterAPIEndpoint
type TwitterAEMap map[string]TwitterAE

type TwitterAPIEndpoint struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	Data   *url.Values
}

type TwitterAPIKey struct {
	Consumer       string
	ConsumerSecret string
	Access         string
	AccessSecret   string
}

type Configuration struct {
	AppName            string                 `json:"appName"`
	AppNetwork         map[string]interface{} `json:"appNetwork"`
	TwitterApiKey      TwitterAPIKey          `json:"twitterApiKey"`
	TwitterApiEndpoint TwitterAEMap           `json:"twitterApiEndpoint"`
	TwitterObservation map[string]string      `json:"twitterObservation"`
}

func (tae *TwitterAE) AddData(key string, value string) {
	if nil == tae.Data {
		tae.Data = &url.Values{}
	}

	tae.Data.Add(key, value)
}

func (tae *TwitterAE) AddDataCollection(data map[string]string) {
	for key, value := range data {
		tae.AddData(key, value)
	}
}
