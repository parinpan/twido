/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T11:22:54+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-16T14:42:45+07:00
 */

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
