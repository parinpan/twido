/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T13:54:55+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-17T00:35:18+07:00
 */

package engine

import (
	"log"
	"strings"
	. "twido/client/redis"
	. "twido/client/twitter"
	. "twido/config"
	. "twido/dataprovider"
)

func CollectMentions() ([]Status, error) {
	redisCacheManager := NewRedisCacheManager()
	lastSavedTweetId, err := redisCacheManager.Get(LastSavedTweetIDKey)
	log.Println("Running CollectMentions with last saved tweet id: " + lastSavedTweetId)

	if nil != err {
		log.Println("Could not get last saved tweet id from redis cache")
	}

	searchingParams := map[string]string{
		"q":                TwidoConfig.TwitterObservation["keyword"],
		"count":            TwidoConfig.TwitterObservation["maxSearchCount"],
		"result_type":      "recent",
		"include_entities": "true",
		"tweet_mode":       "extended",
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

func CollectMentionsAsStatusesToReplyBack() (*StatusesToReplyBack, error) {
	statuses, err := CollectMentions()

	if nil != err {
		log.Println("Couldn't load eligible mentions to reply")
		return nil, err
	}

	/* it's the status that get replied by user which may be containing a video */
	var statusIdToLookup []string

	for _, status := range statuses {
		statusIdToLookup = append(statusIdToLookup, status.InReplyToStatusIDStr)
	}

	if len(statusIdToLookup) == 0 {
		return nil, nil
	}

	service := TwitterService{}
	statusesThatGetReplied, err := service.Statuses.Lookup(map[string]string{
		"id":                   strings.Join(statusIdToLookup, ","),
		"tweet_mode":           "extended",
		"include_entites":      "true",
		"include_ext_alt_text": "true",
	})

	if nil != err {
		log.Println("Couldn't get statuses that get replied by users")
		return nil, err
	}

	var index = 0
	var statusesToReplyBack = &StatusesToReplyBack{}

	for _, status := range *statusesThatGetReplied {
		for _, media := range status.ExtendedEntities.Media {
			if len(media.VideoInfo.Variants) > 0 {
				statusesToReplyBack.AddToQueue(ReplyBody{
					OriginalStatus: statuses[index],
					VideoVariant:   media.VideoInfo.GetHighestQualityVideoVariant(),
				})

				index += 1
			}
		}
	}

	return statusesToReplyBack, nil
}
