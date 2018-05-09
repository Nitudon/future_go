package main

import (
	"net/http"

	"./application/outgame"
	"./domain/service/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/test", outgame.TestRoom)

	websocket.Routing(router)

	http.ListenAndServe(":8080", router)
}
