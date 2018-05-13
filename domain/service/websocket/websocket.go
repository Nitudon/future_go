package websocket

import (
	"../../../application/outgame"
	"../../../application/ingame"
	"../../../infra"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"log"
	"sync"
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
	lock := new(sync.Mutex)
	lock.Lock()

	if len(outgame.Users) < 1 {
		return
	}

	var err error

	outgame.Users[len(outgame.Users)-1].Session = conn
	outgame.Users[len(outgame.Users)-1].Client, err = infra.CreateRedisClient()
	if err != nil{
		log.Print(err.Error())
		return
	}
	outgame.Users[len(outgame.Users)-1].Pubsub = outgame.Users[len(outgame.Users)-1].Client.Subscribe(outgame.ROOM_ID)

	lock.Unlock()
	go ingame.ReceiveSyncData(outgame.Users[len(outgame.Users)-1])
}

func messageHandle(conn *melody.Session, msg []byte)  {
	data := string(msg[:])

	infra.PublishData(outgame.ROOM_ID,data)
}