package twitter

import (
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
		return nil, err
	}

	return resultInterface, err
}
