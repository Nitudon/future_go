package ingame

import (
	"../../domain/data"
	"github.com/gin-gonic/gin"
	_ "gopkg.in/olahol/melody.v1"
)

type client domain.User

func SyncPlayer(g *gin.Context) {

}

func (user *client) receiveSyncData() {
	for {
		if user.Session.IsClosed() {
			break
		}

		if msg, err := user.Pubsub.ReceiveMessage(); err == nil {
			user.Session.Write([]byte(msg.Payload))
		}
	}
}
