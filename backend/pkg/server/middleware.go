package server

import (
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
		if err != nil{
			//respondError(w, http.StatusBadRequest, err)
			//TODO what if there's no cookie
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
			//TODO don't know what to do when authentication is wrong
			respondError(w, http.StatusBadRequest, err)
			return
		}


		//TODO should not exist when joining
		if (exists && r.Method == "GET") || (!exists && r.Method == "POST"){
			h.ServeHTTP(w, r)
		} else {
			//TODO do sth not to go to handler
			return
		}

	})
}