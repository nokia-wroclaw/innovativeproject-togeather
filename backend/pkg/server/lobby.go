package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type lobbyHandler struct {
	lobbyService core.LobbyService
}

func (h *lobbyHandler) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	lobbies, err := h.lobbyService.List(ctx)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, lobbies)
}

type createLobbyRequest struct {
	RestaurantID int       `json:"restaurant_id, required"`
	OwnerName    string       `json:"owner_name, required"`
	Expires      time.Time `json:"expires, required"`
	Address      string    `json:"address, required"`
}

type editLobbyRequest struct {
	RestaurantID int       `json:"restaurant_id, required"`
	OwnerID    	 int       `json:"owner_id, required"`
	Expires      time.Time `json:"expires, required"`
	Address      string    `json:"address, required"`
}

type joinLobbyRequest struct {
	UserName string `json:"user_name"`
}

func (h *lobbyHandler) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := r.Cookie("user-id")
	if cookie != nil {
		respondError(w, http.StatusBadRequest,
			errors.New("given user already belongs to a lobby"))
		return
	}

	var request createLobbyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	lobby, userID, err := h.lobbyService.Create(
		ctx,
		request.RestaurantID,
		request.OwnerName,
		&request.Expires,
		request.Address,
	)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	addCookie(w, "user-id", strconv.Itoa(userID), 24*60*60, "/lobbies")
	respondJSON(w, http.StatusOK, lobby)
}

func (h *lobbyHandler) edit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	lobbyID, err := strconv.Atoi(chi.URLParam(r, "lobbyID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	var request editLobbyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	lobby, err := h.lobbyService.Edit(
		ctx,
		lobbyID,
		request.RestaurantID,
		request.OwnerID,
		&request.Expires,
		request.Address,
	)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, lobby)
}

func (h *lobbyHandler) join(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	lobbyID, err := strconv.Atoi(chi.URLParam(r, "lobbyID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	var request joinLobbyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.lobbyService.Join(ctx, lobbyID, request.UserName)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	addCookie(w, "user-id", strconv.Itoa(user.ID), 24*60*60, "/lobbies")
	respondJSON(w, http.StatusOK, user)
}

func (h *lobbyHandler) get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	lobbyID, err := strconv.Atoi(chi.URLParam(r, "lobbyID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	lobby, err := h.lobbyService.Get(ctx, lobbyID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, lobby)
}
