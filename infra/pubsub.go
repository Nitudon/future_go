package infra

import (
	"errors"
	"log"

	"../domain/data"

	"github.com/go-redis/redis"
)

const (
	Addr = ""
)

type User domain.User

var (
	cl_publish *redis.Client
)

func init() {
	var err error
	cl_publish, err = CreateRedisClient()
	if err != nil {
		log.Println(err.Error())
	}
}

func CreateRedisClient() (*redis.Client, error) {
	cl := redis.NewClient(&redis.Options{
		Addr: Addr,
	})
	if cl == nil {
		err := errors.New("cant create client")
		return nil, err
	}

	return cl, nil
}

func PublishData(room string, data string) error {
	err := cl_publish.Publish(room, data).Err()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (user *User) SubscribeRoom(room string) *User {
	user.Pubsub.Subscribe(room)
	return user
}
