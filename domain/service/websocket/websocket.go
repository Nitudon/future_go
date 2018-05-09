package websocket

import (
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
}

func websocketHandle(conn *melody.Session) {

}
