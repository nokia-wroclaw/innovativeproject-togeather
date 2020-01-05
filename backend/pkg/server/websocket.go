package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type wsHandler struct {
	hub *hub
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func (wsh *wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading %s", err)
		return
	}
	c := &connection{send: make(chan []byte, 256), hub: wsh.hub}
	c.hub.addConnection(c)
	defer c.hub.removeConnection(c)

	var wg sync.WaitGroup
	wg.Add(2)
	go c.writer(&wg, wsConn)
	go c.reader(&wg, wsConn)
	wg.Wait()

	wsConn.Close()
}

type longLatStruct struct {
	Long float64 `json:"longitude"`
	Lat  float64 `json:"latitude"`
}

func (wsh *wsHandler) longLatHandler(w http.ResponseWriter, r *http.Request) {
	var coordinates longLatStruct
	if err := json.NewDecoder(r.Body).Decode(&coordinates); err != nil {
		log.Printf("ERROR: %s", err)
		http.Error(w, "Bad request", http.StatusTeapot)
		return
	}
	defer r.Body.Close()
	go writer(&coordinates)
}




