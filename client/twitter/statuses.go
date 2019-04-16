package twitter

import (
	. "twido/dataprovider"
)

type StatusesService struct {
}

func (ss *StatusesService) Show(params map[string]string) (*Status, error) {
	result, err := LoadService(
		&ServiceBody{SlugName: "show-status", Params: params},
		nil,
	)

	if nil != err {
		return nil, err
	}

	return result.(*Status), err
}

func (ss *StatusesService) Lookup(params map[string]string) ([]Status, error) {
	result, err := LoadService(
		&ServiceBody{SlugName: "lookup-statuses", Params: params},
		nil,
	)

	if nil != err {
		return nil, err
	}

	return result.([]Status), err
}
