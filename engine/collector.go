/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T13:54:55+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-16T14:39:24+07:00
 */

package engine

import (
	"log"
	. "twido/client/redis"
	. "twido/client/twitter"
	. "twido/config"
	. "twido/dataprovider"
)

func collectEligibleMentionsToReply() ([]Status, error) {
	redisCacheManager := NewRedisCacheManager()
	lastSavedTweetId, err := redisCacheManager.Get(LastSavedTweetIDKey)

	if nil != err {
		log.Println("Could not get last saved tweet id from redis cache")
	}

	searchingParams := map[string]string{
		"q":     TwidoConfig.TwitterObservation["keyword"],
		"count": TwidoConfig.TwitterObservation["maxSearchCount"],
	}

	if "" != lastSavedTweetId {
		searchingParams["since_id"] = lastSavedTweetId
	}

	var result *TwitterSearch
	service := TwitterService{}
	result, err = service.Search.Tweets(searchingParams)

	if nil != err {
		log.Println("Couldn't connect to twitter search")
		return nil, err
	}

	return result.Statuses, nil
}
