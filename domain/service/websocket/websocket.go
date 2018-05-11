package websocket

import (
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func Routing(r *gin.Engine) {
	m := melody.New()

	group := r.Group("ws")
	group.GET("/", func(g *gin.Context) {
		m.HandleRequest(g.Writer, g.Request)
	})
	m.HandleConnect(websocketHandle)

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Println(msg)
		m.Broadcast(msg)
	})
}

func websocketHandle(conn *melody.Session) {

}
