package websocket

import (
	"../../../application/ingame"
	"../../../application/outgame"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func Routing(r *gin.Engine) {
	m := melody.New()

	group := r.Group("ws")
	group.GET("/connect", func(g *gin.Context) {
		m.HandleRequest(g.Writer, g.Request)
	})

	group.POST("/syncPlayer", func(g *gin.Context) {
		ingame.SyncPlayer(g)
	})
	m.HandleConnect(connectHandle)
}

func connectHandle(conn *melody.Session) {
	if len(outgame.Users) < 1 {
		return
	}

	user := outgame.Users[len(outgame.Users)-1]
	user.Session = conn
}
