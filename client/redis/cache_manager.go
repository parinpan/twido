/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T13:57:49+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-16T14:13:19+07:00
 */

package redis

import (
	"time"

	"github.com/go-redis/redis"
)

const (
	MinimumTimeSetPersistMock = time.Second * 500
	LastSavedTweetIDKey       = "redis_last_saved_tweet_id"
)

type RedisCacheManager struct {
	Client *redis.Client
}

func NewRedisCacheManager() *RedisCacheManager {
	return &RedisCacheManager{
		Client: NewRedisConnection(),
	}
}

func (rcm *RedisCacheManager) Get(key string) (string, error) {
	stringCmd := rcm.Client.Get(key)

	if nil != stringCmd.Err() {
		return "", stringCmd.Err()
	}

	return stringCmd.Val(), nil
}

func (rcm *RedisCacheManager) Set(key string, value interface{}, expiration time.Duration) bool {
	statusCmd := rcm.Client.Set(key, value, expiration)
	return nil == statusCmd.Err()
}

func (rcm *RedisCacheManager) SetPersist(key string, value interface{}) bool {
	rcm.Set(key, value, MinimumTimeSetPersistMock)
	boolCmd := rcm.Client.Persist(key)
	return boolCmd.Val()
}
