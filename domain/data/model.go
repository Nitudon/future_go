package domain

import (
	"github.com/go-redis/redis"
	"github.com/olahol/melody"
)

type SyncType int

const (
	ROOM SyncType = iota
	OBJECT
	PLAYER
)

type Room struct {
	RoomId  string       `json:"Id"`
	Name    string       `json:"Name"`
	Time    float32      `json:"Time"`
	Players []SyncPlayer `json:"Players"`
}

type User struct {
	Client  *redis.Client   `json:"client"`
	Session *melody.Session `json:"sesssion"`
	Pubsub  *redis.PubSub   `json:"pubsub"`
	Player  *SyncPlayer     `json:"player"`
}

type SyncMessage struct {
	SyncType SyncType `json:"SyncType"`
	Message  string   `json:"Message"`
}

type SyncObject struct {
	Id          string  `json:"Id"`
	X           float32 `json:"PositionX"`
	Y           float32 `json:"PositionY"`
	Z           float32 `json:"PositionZ"`
	IsDestroyed bool    `json:"IsDestroyed"`
}

type SyncPlayer struct {
	Id          string  `json:"Id"`
	PlayerId    int     `json:"PlayerId"`
	Name        string  `json:"Name"`
	X           float32 `json:"PositionX"`
	Y           float32 `json:"PositionY"`
	Z           float32 `json:"PositionZ"`
	IsDestroyed bool    `json:"IsDestroyed"`
	Hp          float32 `json:"Hp"`
}
