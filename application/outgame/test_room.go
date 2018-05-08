package outgame

import (
	"github.com/gin-gonic/gin"

	"../../domain/data"
)

func TestRoom(g *gin.Context) {
	user1 := new(domain.User)

	user1.Player.Name = "test1"
	user1.Player.Id = 1

	user2 := new(domain.User)

	user2.Player.Name = "test2"
	user2.Player.Id = 2

	user3 := new(domain.User)

	user3.Player.Name = "test3"
	user3.Player.Id = 3

	room.Players = []domain.SyncPlayer{
		user1.Player,
		user2.Player,
		user3.Player,
	}

	g.JSON(200, room)

}
