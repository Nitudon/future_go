package websocket

import (
	"../../../application/outgame"
	"../../../application/ingame"
	"../../../infra"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"log"
)

func Routing(r *gin.Engine) {
	m := melody.New()

	group := r.Group("ws")
	group.GET("/connect", func(g *gin.Context) {
		m.HandleRequest(g.Writer, g.Request)
	})

	m.HandleConnect(connectHandle)
	m.HandleMessage(messageHandle)
}

func connectHandle(conn *melody.Session) {
	if len(outgame.Users) < 1 {
		return
	}

	var err error

	user := outgame.Users[len(outgame.Users)-1]
	user.Session = conn
	user.Client, err = infra.CreateRedisClient()
	if err != nil{
		log.Print(err.Error())
		return
	}
	user.Pubsub = user.Client.Subscribe(outgame.ROOM_ID)
	go ingame.ReceiveSyncData(user)
}

func messageHandle(conn *melody.Session, msg []byte)  {
	data := string(msg[:])

	infra.PublishData(outgame.ROOM_ID,data)
}