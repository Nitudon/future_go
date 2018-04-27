package domain

import (
	"github.com/go-redis/redis"
)

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
	Position Vector `json"position"`
}

type Sync_player struct {
	Sync_object,
	Id int16 `json:"id"`
	Hp int16 `json`
}
