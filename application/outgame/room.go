package outgame

import (
	"container/list"

	"../../domain/data"
)

var (
	room  domain.Room
	users *list.List
)

const ROOM_ID = "Room1"

func init() {
	room = domain.Room{Id: ROOM_ID, Name: ROOM_ID}
	users = list.New()
}

func CreateRoom(id string, owner *domain.User) {
	room.Name = id
	users.PushBack(owner)
}

func JoinRoom(user domain.User) {
	users.PushBack(user)
}

func LeftRoom(user domain.User) {
	users.Remove(&list.Element{Value: user})
}
