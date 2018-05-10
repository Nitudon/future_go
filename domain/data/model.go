package domain

import (
	"github.com/go-redis/redis"
)

type Room struct {
	Id      string       `json:"Id"`
	Name    string       `json:"Name"`
	Players []SyncPlayer `json:"Players"`
}

type User struct {
	Client *redis.Client `json:"client"`
	Pubsub *redis.PubSub `json:"pubsub"`
	Player SyncPlayer    `json:"player"`
}

type SyncObject struct {
	Id int     `json:"Id"`
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
	Z  float32 `json:"z"`
}

type SyncPlayer struct {
	Id   int     `json:"Id"`
	Name string  `json:"Name"`
	Hp   float32 `json:"Hp"`
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	Z    float32 `json:"z"`
}
