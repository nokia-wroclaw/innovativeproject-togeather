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
	Expires      time.Time `json:"expires, required"`
	Address      string    `json:"address, required"`
}

type orderRequest struct {
	MealID int `json:"meal_id"`
}

//type editLobbyRequest struct {
//	RestaurantID int       `json:"restaurant_id, required"`
//	OwnerID    	 int       `json:"owner_id, required"`
//	Expires      time.Time `json:"expires, required"`
//	Address      string    `json:"address, required"`
//}

func (h *lobbyHandler) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value(UserKey).(*core.User)
	if user == nil {
		respondError(w, http.StatusBadRequest,
			errors.New("join lobby: unauthorized"))
	}

	var request createLobbyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	lobby, _, err := h.lobbyService.Create(
		ctx,
		request.RestaurantID,
		user.Name,
		&request.Expires,
		request.Address,
	)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	//addCookie(w, "user-id", strconv.Itoa(userID), 24*60*60, "/api/lobbies")
	respondJSON(w, http.StatusOK, lobby)
}

//func (h *lobbyHandler) edit(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//
//	lobbyID, err := strconv.Atoi(chi.URLParam(r, "lobbyID"))
//	if err != nil {
//		respondError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	var request editLobbyRequest
//	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//		respondError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	lobby, err := h.lobbyService.Edit(
//		ctx,
//		lobbyID,
//		request.RestaurantID,
//		request.OwnerID,
//		&request.Expires,
//		request.Address,
//	)
//	if err != nil {
//		respondError(w, http.StatusBadRequest, err)
//		return
//	}
//
//	respondJSON(w, http.StatusOK, lobby)
//}

func (h *lobbyHandler) join(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value(UserKey).(*core.User)
	if user == nil {
		respondError(w, http.StatusBadRequest,
			errors.New("join lobby: unauthorized"))
	}

	lobbyID, err := strconv.Atoi(chi.URLParam(r, "lobbyID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	//var request joinLobbyRequest
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	respondError(w, http.StatusBadRequest, err)
	//	return
	//}

	err = h.lobbyService.Join(ctx, lobbyID, user.Name)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	//addCookie(w, "user-id", strconv.Itoa(user.ID), 24*60*60, "/api/lobbies")
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

func (h *lobbyHandler) addToCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value(UserKey).(*core.User)
	if user == nil {
		respondError(w, http.StatusBadRequest,
			errors.New("add to cart: unauthorized"))
		return
	}

	lobbyID, err := strconv.Atoi(chi.URLParam(r, "lobbyID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	var request orderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	err = h.lobbyService.AddToCart(ctx, user.ID, lobbyID, request.MealID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, nil)
}

func (h *lobbyHandler) delFromCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value(UserKey).(*core.User)
	if user == nil {
		respondError(w, http.StatusBadRequest,
			errors.New("delete from cart: unauthorized"))
		return
	}

	lobbyID, err := strconv.Atoi(chi.URLParam(r, "lobbyID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	var request orderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	err = h.lobbyService.DelFromCart(ctx, user.ID, lobbyID, request.MealID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, nil)
}
