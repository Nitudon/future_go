package domain

import (
	"github.com/go-redis/redis"
)

type Room struct {
	Id      string       `json:"id"`
	Name    string       `json:"name"`
	Players []SyncPlayer `json:"players"`
}

type User struct {
	Client *redis.Client `json:"client"`
	Pubsub *redis.PubSub `json:"pubsub"`
	Player SyncPlayer    `json:"player"`
}

type Vector struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type SyncObject struct {
	Id int16   `json:"id"`
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
	Z  float32 `json:"z"`
}

type SyncPlayer struct {
	Id   int16   `json:"id"`
	Name string  `json:"name"`
	Hp   int16   `json:"hp"`
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	Z    float32 `json:"z"`
}
