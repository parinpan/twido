/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T10:31:05+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   nakama
 * @Last modified time: 2019-04-16T22:47:35+07:00
 */

package dataprovider

import (
	"net/url"
)

type TwitterAE TwitterAPIEndpoint
type TwitterAEMap map[string]TwitterAE

type RedisConnection struct {
	Addr     string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

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

type Rebrandly struct {
	BaseUrlApi string `json:"baseUrlApi"`
	ApiKey     string `json:"apiKey"`
	Domain     string `json:"domain"`
	Active     bool   `json:"active"`
}

type UrlShortener struct {
	Rebrandly Rebrandly `json:"rebrandly"`
}

type Configuration struct {
	AppName            string                 `json:"appName"`
	AppNetwork         map[string]interface{} `json:"appNetwork"`
	RedisConnection    RedisConnection        `json:"redisConnection"`
	TwitterApiKey      TwitterAPIKey          `json:"twitterApiKey"`
	TwitterApiEndpoint TwitterAEMap           `json:"twitterApiEndpoint"`
	TwitterObservation map[string]string      `json:"twitterObservation"`
	UrlShortener       UrlShortener           `json:"urlShortener"`
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
