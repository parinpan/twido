/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T12:16:16+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-17T11:49:05+07:00
 */

package twitter

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	. "twido/client"
	. "twido/config"
	. "twido/dataprovider"
)

type TwitterClientRequest struct {
	Endpoint TwitterAE
	Client   *http.Client
	Request  *http.Request
}

type TwitterOauth struct {
	TwitterClientReq *TwitterClientRequest
}

/* Collection of TwitterClientRequest method implementation */

func NewTwitterClientRequest(endpoint TwitterAE) *TwitterClientRequest {
	var networkTimeout = time.Duration(
		int64(TwidoConfig.AppNetwork["timeout"].(float64)),
	)

	client := &http.Client{
		Timeout: time.Second * networkTimeout,
	}

	return &TwitterClientRequest{
		Client:   client,
		Endpoint: endpoint,
	}
}

func (tcr *TwitterClientRequest) request() (*http.Response, error) {
	if !IsTwitterCredentialSet() {
		return nil, errors.New("Twitter Credential isn't set in this app yet.")
	}

	var err error
	endpoint := tcr.Endpoint
	method := strings.ToUpper(tcr.Endpoint.Method)

	switch method {
	case "GET":
		tcr.Request, err = tcr.get()
	case "POST":
		tcr.Request, err = tcr.post()
	}

	if nil != err {
		log.Println("Couldn't connect to ", endpoint.URL, " endpoint")
		return nil, err
	}

	to := NewTwitterOauth(tcr)
	oauthHeader, err := to.getOauthHeader()

	if nil != err {
		log.Println("Couldn't create oauth header")
		return nil, err
	}

	tcr.Request.Header.Add("Authorization", oauthHeader)
	response, err := tcr.Client.Do(tcr.Request)

	if nil != err {
		log.Println("Couldn't get request's response to ", endpoint.URL)
		return nil, err
	}

	return response, nil
}

func (tcr *TwitterClientRequest) get() (*http.Request, error) {
	req, err := http.NewRequest(tcr.Endpoint.Method, tcr.Endpoint.URL, nil)

	if nil != err {
		return nil, err
	}

	req.URL.RawQuery = tcr.Endpoint.Data.Encode()

	return req, nil
}

/* This function is still considered having strange behavior, fix it later */
func (tcr *TwitterClientRequest) post() (*http.Request, error) {
	data := strings.NewReader(tcr.Endpoint.Data.Encode())
	req, err := http.NewRequest(tcr.Endpoint.Method, tcr.Endpoint.URL, data)

	if nil != err {
		return nil, err
	}

	req.URL.RawQuery = tcr.Endpoint.Data.Encode()

	return req, nil
}

func (tcr *TwitterClientRequest) Raw() ([]byte, error) {
	resp, err := tcr.request()

	if nil != err {
		return nil, err
	}

	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		return nil, err
	}

	return responseData, nil
}

func (tcr *TwitterClientRequest) Result(convertName string) (interface{}, error) {
	raw, err := tcr.Raw()

	if nil != err {
		return nil, err
	}

	converter := NewRequestConverter(
		convertName,
		raw,
	)

	result, err := converter.Convert()

	if nil != err {
		log.Println("Could not convert raw data to object result")
		return nil, err
	}

	return result, nil
}

/* End of Collection of TwitterClientRequest method implementation */

/* Collection of TwitterOauthV1 method implementation */

func NewTwitterOauth(tcr *TwitterClientRequest) *TwitterOauth {
	return &TwitterOauth{TwitterClientReq: tcr}
}

func (to TwitterOauth) getTimestamp() string {
	t := time.Now()
	return strconv.FormatInt(t.Unix(), 10)
}

func (to TwitterOauth) getUniqueNonce(length int) (string, error) {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b), nil
}

func (to TwitterOauth) getSignature(params map[string]string) string {
	return SignatureBase(
		to.TwitterClientReq.Request,
		params,
	)
}

func (to TwitterOauth) getSignedSignature(oauthParams map[string]string) (string, error) {
	req := to.TwitterClientReq.Request
	params, err := CollectParameters(req, oauthParams)

	if nil != err {
		return "", nil
	}

	signature := SignatureBase(req, params)
	signingKey := strings.Join(
		[]string{
			TwidoConfig.TwitterApiKey.ConsumerSecret,
			TwidoConfig.TwitterApiKey.AccessSecret,
		},
		"&",
	)

	mac := hmac.New(
		sha1.New,
		[]byte(signingKey),
	)

	if _, err := mac.Write([]byte(signature)); nil != err {
		return "", err
	}

	signedSignature := base64.StdEncoding.EncodeToString(
		mac.Sum(nil),
	)

	return PercentEncode(signedSignature), nil
}

func (to TwitterOauth) getOauthHeader() (string, error) {
	currentTimeStamp := to.getTimestamp()
	uniqueNonce, err := to.getUniqueNonce(32)
	oauthData := []string{}

	if nil != err {
		log.Println("Could not get nonce for app's authentication")
		return "", err
	}

	oauthParams := map[string]string{
		"oauth_consumer_key":     TwidoConfig.TwitterApiKey.Consumer,
		"oauth_nonce":            uniqueNonce,
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        currentTimeStamp,
		"oauth_token":            TwidoConfig.TwitterApiKey.Access,
		"oauth_version":          "1.0",
	}

	signedSignature, err := to.getSignedSignature(oauthParams)

	if nil != err {
		log.Println("Could not create signed signature for the app.")
		return "", err
	}

	oauthParams["oauth_signature"] = signedSignature

	for key, value := range oauthParams {
		oauthData = append(oauthData, key+"="+value)
	}

	return "OAuth " + strings.Join(oauthData, ","), nil
}

/* End of Collection of TwitterOauthV1 method implementation */
