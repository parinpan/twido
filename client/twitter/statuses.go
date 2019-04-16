/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T12:16:52+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-17T00:35:05+07:00
 */

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

func (ss *StatusesService) Lookup(params map[string]string) (*[]Status, error) {
	result, err := LoadService(
		&ServiceBody{SlugName: "lookup-statuses", Params: params},
		nil,
	)

	if nil != err {
		return nil, err
	}

	return result.(*[]Status), err
}

func (ss *StatusesService) Update(params map[string]string) (*Status, error) {
	result, err := LoadService(
		&ServiceBody{SlugName: "update-status", Params: params},
		nil,
	)

	if nil != err {
		return nil, err
	}

	return result.(*Status), err
}
