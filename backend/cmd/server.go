package main

import (
	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/server"
	"log"
	"net/http"
)

func runServer() {
	srvr := server.New()

	if err := http.ListenAndServe(":8000", srvr); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}

// e.g. methods to connect to db, redis go here