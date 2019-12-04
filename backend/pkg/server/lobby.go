package server

import (
	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
	"net/http"
)

type lobbyHandler struct {
	lobbyService core.LobbyService
}

func (h *lobbyHandler) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

}
