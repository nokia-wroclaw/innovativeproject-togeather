package main

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/server"
)

func runServer() {
	srvr := server.New()

	if err := http.ListenAndServe(":8000", srvr); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}

// e.g. methods to connect to db, redis go here
func dbConnect() (*sqlx.DB, error) {
	connectionString := "user=postgres dbname=postgres password=postgres" +
		" sslmode=disable"

	db, err := sqlx.Connect("postgres", connectionString)
	return db, err
}

func redisConnect() (*redis.Client, error) {
	//address := os.Getenv("REDIS_URL")

	address := "eather-redis:6379"
	client := redis.NewClient(
		&redis.Options{
			Addr: address,
		})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
