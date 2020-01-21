package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type lobbyMiddleware struct {
	lobbyService core.LobbyService
}

func (m *lobbyMiddleware) cookiesMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user-id")
		if err != nil && r.Method == "GET"{
			respondError(w, http.StatusBadRequest, errors.New("user not authorized"))
			return
		} else if err != nil{
			h.ServeHTTP(w, r)
			return
		}

		clientID, err := strconv.Atoi(cookie.Value)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		lobbyID, err := strconv.Atoi(chi.URLParam(r, "lobbyID"))
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		ctx := r.Context()
		exists, err := m.lobbyService.BelongsToLobby(ctx, clientID, lobbyID)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		if (exists && r.Method == "GET") || (!exists && r.Method == "POST"){
			h.ServeHTTP(w, r)
		} else if exists && r.Method == "POST" {
			respondError(w, http.StatusBadRequest, errors.New("given user already belongs to this lobby"))
			return

		} else {
			respondError(w, http.StatusBadRequest, errors.New("given user does not belong to this lobby"))
			return
		}

	})
}