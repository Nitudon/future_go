package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func Routing() {
	router := gin.Default()
	m := melody.New()

	group := router.Group("ws")
	group.GET("/", func(g *gin.Context) {
		m.HandleRequest(g.Writer, g.Request)
	})
	m.HandleConnect(websocketHandle)

	http.ListenAndServe(":3000", router)
}

func websocketHandle(conn *melody.Session) {

}
