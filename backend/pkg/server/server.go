package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type Server struct {
	// smth like config goes here
	router chi.Router
}

func New(
	restaurantService core.RestaurantService,
	lobbyService core.LobbyService,
) *Server {
	s := &Server{}

	r := chi.NewRouter()

	//h := newHub()

	r.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.StripSlashes,
		middleware.Timeout(10*time.Second),
	)

	// add handlers
	//webSocketHandler := wsHandler{hub: h}

	pingHandler := pingHandler{}
	restaurantHandler := restaurantHandler{restaurantService:restaurantService}
	lobbyHandler := lobbyHandler{lobbyService: lobbyService}

	hub := newHub()
	go hub.run()

	r.HandleFunc("/", wsEndpoint)
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	r.Route("/api", func(r chi.Router) {
		//r.HandleFunc("/ws", webSocketHandler.ServeHTTP)
		//r.Route("/ws", func(r chi.Router) {
		//	r.Get("/", webSocketHandler.ServeHTTP)
		//})
		//router := mux.NewRouter()
		//r.HandleFunc("/longlat", webSocketHandler.longLatHandler)
		//router.HandleFunc("/", rootHandler).Methods("GET")
		//router.HandleFunc("/ws", wsHandler)
		//go echo()

		r.Route("/restaurants", func(r chi.Router) {
			r.Get("/", restaurantHandler.listRestaurants)
			r.Route("/{restaurantID}", func(r chi.Router){
				r.Get("/", restaurantHandler.getRestaurant)
				r.Get("/menu", restaurantHandler.restaurantMenu)
			})
		})

		r.Route("/lobbies", func(r chi.Router) {
			r.Get("/", lobbyHandler.list)
			r.Post("/", lobbyHandler.create)
		})

		r.Route("/ping", func(r chi.Router) {
			r.Get("/", pingHandler.ping)
		})
	})

	s.router = r

	return s
}

//var upgrader = websocket.Upgrader{
//	ReadBufferSize:    1024,
//	WriteBufferSize:   1024,
//}
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println(r.URL)
	http.ServeFile(w, r, "chat.html")
}

//func wsEndpoint(w http.ResponseWriter, r *http.Request) {
//	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
//
//	ws, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Println(err)
//	}
//
//	log.Println("CLIENT CONNECTED")
//	err = ws.WriteMessage(1, []byte("Hi Client!"))
//	if err != nil {
//		log.Println(err)
//	}
//
//	reader(ws)
//}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

//func reader(conn *websocket.Conn) {
//	for {
//		read in a message
		//messageType, p, err := conn.ReadMessage()
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//print out that message for clarity
		//fmt.Println(string(p))
		//
		//if err := conn.WriteMessage(messageType, p); err != nil {
		//	log.Println(err)
		//	return
		//}
	//
	//}
//}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s.router.ServeHTTP(w, r)
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(jsonData)
}

func respondError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
