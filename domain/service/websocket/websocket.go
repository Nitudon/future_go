package websocket

import (
	"net/http"

	"github.com/trevex/golem"
)

func Routing() {
	router := golem.NewRouter()
	router.On("connect", websocketHandle)

	http.HandleFunc("/connect", router.Handler())
	http.ListenAndServe(":8080", nil)
}

func websocketHandle(conn *golem.Connection) {

}
