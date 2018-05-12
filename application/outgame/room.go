package outgame

import (
	"../../domain/data"
	"../../infra"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var (
	room  domain.Room
	Users []*domain.User
)

const ROOM_ID = "Room1"

func init() {
	room = domain.Room{RoomId: ROOM_ID, Name: ROOM_ID}
	Users = []*domain.User{}
}

func InitRoom(group *gin.RouterGroup) {
	group.GET("/join", JoinRoom)
}

func JoinRoom(g *gin.Context) {
	var err error

	name := "test"

	user := new(domain.SyncPlayer)
	user.PlayerId = len(room.Players)
	user.Name = name
	user.Id = xid.New().String()

	cl := new(domain.User)
	cl.Client, err = infra.CreateRedisClient()
	if err != nil {
		g.JSON(
			500, gin.H{
				"error": "Redis Session Error",
			},
		)
	}
	cl.Player = user

	room.Players = append(room.Players, *user)
	g.JSON(
		200, room,
	)
}

func createRoom(id string, owner domain.SyncPlayer) {
	room.Name = id
	room.Players = append(room.Players, owner)
}
