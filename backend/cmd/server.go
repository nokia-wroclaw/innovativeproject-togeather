package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/server"
)

func runServer() {
	srvr := server.New()

	if err := http.ListenAndServe(":8000", srvr); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}

func redisConnect() (*redis.Client, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr: ":6379",
		})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}


// e.g. methods to connect to db, redis go here