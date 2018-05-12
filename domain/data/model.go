package domain

import (
	"github.com/go-redis/redis"
)

type SyncType int

const (
	ROOM SyncType = iota
	OBJECT
	PLAYER
)

type Room struct {
	Id      string        `json:"room_id"`
	Name    string        `json:"name"`
	Players []*SyncPlayer `json:"players"`
}

type User struct {
	Client *redis.Client `json:"client"`
	Pubsub *redis.PubSub `json:"pubsub"`
	Player SyncPlayer    `json:"player"`
}

type SyncMessage struct {
	SyncType SyncType `json:"sync_type"`
	Message  string   `json:"message"`
}

type SyncObject struct {
	Id          string  `json:"id"`
	X           float32 `json:"position_x"`
	Y           float32 `json:"position_y"`
	Z           float32 `json:"position_z"`
	IsDestroyed bool    `json:"is_destroyed"`
}

type SyncPlayer struct {
	Id          string  `json:"id"`
	PlayerId    int     `json:"player_id`
	Name        string  `json:"name"`
	X           float32 `json:"position_x"`
	Y           float32 `json:"position_y"`
	Z           float32 `json:"position_z"`
	IsDestroyed bool    `json:"is_destroyed"`
	Hp          float32 `json:"hp"`
}
