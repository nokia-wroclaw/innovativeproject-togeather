package server

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
	"net/http"
	"strconv"
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
	Address string `json:"address, required"`
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
		request.Address,
	)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, lobby)
}

func (h *lobbyHandler) edit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	lobbyID, err := strconv.Atoi(chi.URLParam(r, "lobbyID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	var request createLobbyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	lobby, err := h.lobbyService.Edit(
		ctx,
		lobbyID,
		request.RestaurantID,
		request.Owner,
		&request.Expires,
		request.Address,
	)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, lobby)
}
