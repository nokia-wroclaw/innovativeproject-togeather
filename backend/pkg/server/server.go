package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

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

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Content-Length", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.StripSlashes,
		middleware.Timeout(10*time.Second),
		cors.Handler,
	)

	// add handlers
	pingHandler := pingHandler{}
	restaurantHandler := restaurantHandler{restaurantService: restaurantService}
	lobbyHandler := lobbyHandler{lobbyService: lobbyService}
	userHandler := userHandler{userService: userService}
	lobbyMiddleware := lobbyMiddleware{lobbyService: lobbyService, userService: userService}
	authHandler := authHandler{userService: userService}

	r.Route("/api", func(r chi.Router) {
		r.Route("/restaurants", func(r chi.Router) {
			r.Get("/", restaurantHandler.list)
			r.Route("/{restaurantID}", func(r chi.Router){
				r.Get("/", restaurantHandler.get)
				r.Get("/menu", restaurantHandler.menu)
			})
		})

		r.Route("/lobbies", func(r chi.Router) {
			r.Get("/", lobbyHandler.list)
			r.Post("/", authMiddleware(lobbyHandler.create, lobbyMiddleware))
			r.Route("/{lobbyID}", func(r chi.Router){
				//r.Put("/", authMiddleware(lobbyHandler.edit, lobbyMiddleware))
				r.Post("/", authMiddleware(lobbyHandler.join, lobbyMiddleware))
				r.Get("/", authMiddleware(lobbyHandler.get, lobbyMiddleware))
				r.Route("/order", func(r chi.Router){
					r.Post("/", authMiddleware(lobbyHandler.addToCart, lobbyMiddleware))
					r.Delete("/", authMiddleware(lobbyHandler.delFromCart, lobbyMiddleware))
					r.Get("/", authMiddleware(lobbyHandler.getCart, lobbyMiddleware))
				})
			})
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.list)
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/", userHandler.get)
		})

		r.Route("/auth", func (r chi.Router) {
			r.Post("/register", authHandler.register)
			r.Post("/login", authHandler.login)
			r.Delete("/logout", authMiddleware(authHandler.logout, lobbyMiddleware))
		})

		r.Route("/ping", func(r chi.Router) {
			r.Get("/", authMiddleware(pingHandler.ping, lobbyMiddleware))
		})
	})

	s.router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func addCookie(w http.ResponseWriter, name string, value string, maxAge int, path string) {
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		MaxAge: maxAge,
		Path: path,
	}
	http.SetCookie(w, &cookie)
}
