package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/server"
)

func runServer(
	restaurantService core.RestaurantService,
) {
	srvr := server.New(
		restaurantService,
	)

	if err := http.ListenAndServe(":8000", srvr); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}

// e.g. methods to connect to db, redis go here
func dbConnect() (*sqlx.DB, error) {
	connectionString := "host=eather-postgres user=postgres " +
		"dbname=postgres password=postgres sslmode=disable"

	db, err := sqlx.Connect("postgres", connectionString)
	return db, err
}

func redisConnect() (*redis.Client, error) {
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
