package ingame

import (
	"../../domain/data"
	_ "gopkg.in/olahol/melody.v1"
	"log"
)

func ReceiveSyncData(user *domain.User) {
	for {
		if msg, err := user.Pubsub.ReceiveMessage(); err == nil {
			log.Println(msg.Payload)
			user.Session.Write([]byte(msg.Payload))
		}else {
			log.Println(err.Error())
		}
	}
}
