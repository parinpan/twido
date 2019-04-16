package twitter

import (
	. "twido/dataprovider"
)

type SearchService struct {
}

func (ss *SearchService) Tweets(params map[string]string) (*TwitterSearch, error) {
	result, err := LoadService(
		&ServiceBody{SlugName: "search-tweets", Params: params},
		nil,
	)

	if nil != err {
		return nil, err
	}

	return result.(*TwitterSearch), err
}
