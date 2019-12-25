package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type Server struct {
	// smth like config goes here
	router chi.Router
}

func New(
	restaurantService core.RestaurantService,
	lobbyService core.LobbyService,
	userService core.UserService,
) *Server {
	s := &Server{}

	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.StripSlashes,
		middleware.Timeout(10*time.Second),
	)

	// add handlers
	pingHandler := pingHandler{}
	restaurantHandler := restaurantHandler{restaurantService: restaurantService}
	lobbyHandler := lobbyHandler{lobbyService: lobbyService}
	userHandler := userHandler{userService: userService}

	r.Route("/api", func(r chi.Router) {
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

		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.listUsers)
		})

		r.Route("/ping", func(r chi.Router) {
			r.Get("/", pingHandler.ping)
		})
	})

	s.router = r

	return s
}

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
