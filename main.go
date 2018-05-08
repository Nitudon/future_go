package main

import (
	"net/http"

	"./application/outgame"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/test", outgame.TestRoom)

	http.ListenAndServe(":8080", router)
}
