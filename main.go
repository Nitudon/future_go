package main

import (
	"net/http"

	"./application/outgame"
	"./domain/service/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	websocket.Routing(router)
	room := router.Group("/room")

	outgame.InitRoom(room)

	http.ListenAndServe(":3000", router)
}
