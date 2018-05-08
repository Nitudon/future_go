package domain

import (
	"github.com/go-redis/redis"
)

type Room struct {
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	Players []Sync_player `json:"players"`
}

type Client struct {
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
	Id int16   `json:"id"`
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
	Z  float32 `json:"z"`
}

type Sync_player struct {
	Id   int16   `json:"id"`
	Name string  `json:"name"`
	Hp   int16   `json:"hp"`
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	Z    float32 `json:"z"`
}
