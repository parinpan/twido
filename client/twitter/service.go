/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T12:13:54+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-17T11:50:18+07:00
 */

package twitter

import (
	"log"
	. "twido/config"
)

type ServiceMiddleware func(tcr *TwitterClientRequest)

type ServiceBody struct {
	SlugName string
	Params   map[string]string
}

type TwitterService struct {
	Search   SearchService
	Timeline TimelineService
	Statuses StatusesService
}

func LoadService(body *ServiceBody, middleware ServiceMiddleware) (interface{}, error) {
	endpoint := TwidoConfig.TwitterApiEndpoint[body.SlugName]
	endpoint.AddDataCollection(body.Params)
	request := NewTwitterClientRequest(endpoint)

	if nil != middleware {
		middleware(request)
	}

	resultInterface, err := request.Result(body.SlugName)

	if nil != err {
		log.Println(err.Error())
		return nil, err
	}

	return resultInterface, err
}
