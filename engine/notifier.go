/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T16:14:02+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   nakama
 * @Last modified time: 2019-04-16T22:42:55+07:00
 */

package engine

import (
	"fmt"
	"log"
	"sync"
	"time"
	. "twido/client/redis"
	. "twido/client/twitter"
	. "twido/config"
	. "twido/dataprovider"
)

func BuildNotificationString(words ...string) string {
	notifStr := fmt.Sprintf(
		"Hi @%s, here it is your video download link %s. ",
		words[0],
		words[1],
	)

	notifStr += fmt.Sprintf(
		"\n\nHope you enjoy it and thank you for using %s ;)",
		words[2],
	)

	return notifStr
}

func NotifyUserTheVideoDownloadLink(strb *StatusesToReplyBack) {
	dateTime := time.Now().String()
	log.Println("NotifyUserTheVideoDownloadLink starts at: " + dateTime)

	if nil == strb {
		log.Println("There's no queue to process. Apps will exit now.")
		return
	}

	maxStatusID := ""
	replyBackQueue := strb.Queue
	redisCacheManager := NewRedisCacheManager()

	var wg sync.WaitGroup
	var replyToIds sync.Map

	for _, replyBody := range replyBackQueue {
		// add the go routine anonymous function to the wait group
		wg.Add(1)

		go func(rb ReplyBody) {
			defer wg.Done()
			user := rb.OriginalStatus.User
			replyToID := rb.OriginalStatus.IDStr
			videoDownloadLink, err := ShortenURLByRebrandly(rb.VideoVariant.URL)

			notificationString := BuildNotificationString(
				user.ScreenName,
				videoDownloadLink,
				TwidoConfig.AppName,
			)

			service := TwitterService{}
			status, err := service.Statuses.Update(map[string]string{
				"in_reply_to_status_id": replyToID,
				"status":                notificationString,
			})

			/*
			   TODO: Create a fallback mechanism when the function faces
			   failure.
			*/
			if nil != err || nil == status {
				log.Println("Couldn't notify @"+user.ScreenName, "about the video link")
				return
			}

			replyToIds.Store(replyToID, true)
		}(replyBody)
	}

	// wait all the marathon runners ;)
	wg.Wait()

	// finding the newest status ID
	for _, replyBody := range replyBackQueue {
		id := replyBody.OriginalStatus.IDStr
		_, isIDRegisteredToReply := replyToIds.Load(id)

		if id > maxStatusID && isIDRegisteredToReply {
			maxStatusID = id
		}
	}

	// save the newest status ID for the next iteration
	if "" != maxStatusID && maxStatusID > "0" {
		oneYearDuration := time.Hour * 24 * 365

		maxStatusIDSet := redisCacheManager.Set(
			LastSavedTweetIDKey,
			maxStatusID,
			oneYearDuration,
		)

		if maxStatusIDSet {
			log.Println("maxStatusID has been set to ID: " + maxStatusID)
		}
	}
}
