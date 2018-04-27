package infra

import(
"../domain/data"

"github.com/go-redis/redis"
)

const(
	Addr = ""
)

func CreateRedisClient() *redis.Client error{
	cl := redis.NewClient(&redis.Options{
		Addr: Addr
	})
} 