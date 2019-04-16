/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T21:09:13+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-17T00:35:20+07:00
 */

package engine

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	. "twido/config"
)

func ShortenURLByRebrandly(longUrl string) (string, error) {
	if !TwidoConfig.UrlShortener.Rebrandly.Active {
		return longUrl, errors.New("Rebrandly service is disable")
	}

	shortenedUrl := longUrl
	var result map[string]interface{}

	jsonData := map[string]interface{}{
		"destination": longUrl,
		"domain": map[string]string{
			"fullName": TwidoConfig.UrlShortener.Rebrandly.Domain,
		},
	}

	jsonStr, _ := json.Marshal(jsonData)
	data := bytes.NewBuffer(jsonStr)

	rebrandlyBaseUrlApi := TwidoConfig.UrlShortener.Rebrandly.BaseUrlApi
	req, err := http.NewRequest("POST", rebrandlyBaseUrlApi, data)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("apikey", TwidoConfig.UrlShortener.Rebrandly.ApiKey)

	if nil != err {
		log.Println("Couldn't connect to rebrandly service")
		return longUrl, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if nil != err {
		log.Println("Couldn't get rebrandly service response")
		return longUrl, nil
	}

	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		log.Println("Couldn't read rebrandly service response")
		return longUrl, nil
	}

	if err := json.Unmarshal(raw, &result); nil != err {
		log.Println("Couldn't parse rebrandly service response")
		return longUrl, nil
	}

	if _, keyExist := result["shortUrl"]; !keyExist {
		log.Println("Couldn't get shortened url from rebrandly service")
		return longUrl, nil
	}

	shortenedUrl = result["shortUrl"].(string)
	shortenedUrl = "https://" + shortenedUrl

	return shortenedUrl, nil
}
