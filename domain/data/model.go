package domain

import (
	"github.com/go-redis/redis"
)

type Room struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Client *redis.Client `json:"client"`
	Pubsub *redis.PubSub `json:"pubsub"`
	Player Sync_player   `json:"player"`
}

type Vector struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type Sync_object struct {
	Id int16 `json:"id"`
	Vector
}

type Sync_player struct {
	Sync_object,
	Hp int16 `json:"hp"`
}
