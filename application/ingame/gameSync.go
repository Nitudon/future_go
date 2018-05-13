package ingame

import (
	"../../domain/data"
	_ "gopkg.in/olahol/melody.v1"
)

func ReceiveSyncData(user *domain.User) {
	for {
		if user.Session.IsClosed() {
			break
		}

		if msg, err := user.Pubsub.ReceiveMessage(); err == nil {
			user.Session.Write([]byte(msg.Payload))
		}
	}
}
