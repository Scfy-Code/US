package sys

import "github.com/go-redis/redis"

var (
	// Session session
	Session session
)

type session struct {
	redisClient redis.UniversalClient
}
