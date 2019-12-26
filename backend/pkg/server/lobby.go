package server

import (
	"encoding/json"
	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
	"net/http"
	"time"
)

type lobbyHandler struct {
	lobbyService core.LobbyService
}

func (h *lobbyHandler) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := h.lobbyService.List(ctx)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, resp)
}

type createLobbyRequest struct {
	RestaurantID int `json:"restaurant_id, required"`
	Owner int `json:"owner, required"`
	Expires time.Time `json:"expires, required"`
	GeoLat float64 `json:"lat, required"`
	GeoLon float64 `json:"lon, required"`
}

func (h *lobbyHandler) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request createLobbyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	lobby, err := h.lobbyService.Create(
		ctx,
		request.RestaurantID,
		request.Owner,
		&request.Expires,
		request.GeoLat,
		request.GeoLon,
	)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, lobby)
}
