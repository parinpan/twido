/**
 * @Author: Fachrin Aulia Nasution <fachrinfan>
 * @Date:   2019-04-16T13:33:56+07:00
 * @Email:  fachrinfan@gmail.com
 * @Last modified by:   fachrinfan
 * @Last modified time: 2019-04-17T00:34:58+07:00
 */

package redis

import (
	. "twido/config"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func NewRedisConnection() *redis.Client {
	if nil == redisClient {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     TwidoConfig.RedisConnection.Addr,
			Password: TwidoConfig.RedisConnection.Password,
			DB:       TwidoConfig.RedisConnection.DB,
		})
	}

	return redisClient
}
