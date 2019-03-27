package database

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Rclient *redis.Client

func NewClient() {
	Rclient := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.79:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := Rclient.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
