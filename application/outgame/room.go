package outgame

import (
	"container/list"

	"../../domain/data"

	"github.com/gin-gonic/gin"
)

var (
	room  domain.Room
	users *list.List
)

const ROOM_ID = "Room1"

func init() {
	room = domain.Room{Id: ROOM_ID, Name: ROOM_ID}
}

func InitRoom(group *gin.RouterGroup) {
	group.GET("/join", JoinRoom)
}

func JoinRoom(g *gin.Context) {
	name := g.Query("name")

	user := new(domain.SyncPlayer)
	user.Id = len(room.Players)
	user.Name = name

	room.Players = append(room.Players, *user)

	joinRoom(*user)
	g.JSON(
		200, room,
	)
}

func createRoom(id string, owner domain.SyncPlayer) {
	room.Name = id
	room.Players = append(room.Players, owner)
}

func joinRoom(user domain.SyncPlayer) {
	room.Players = append(room.Players, user)
}

func leftRoom(user domain.SyncPlayer) {
	//users.Remove(&list.Element{Value: user})
}
