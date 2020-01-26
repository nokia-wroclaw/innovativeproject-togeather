package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type authHandler struct {
	userService core.UserService
}

type registerUser struct {
	Name string `json:"name,required"`
}

func (h *authHandler) register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload registerUser
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	userID, err := h.userService.Create(ctx, payload.Name)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusCreated, userID)
}

type loginUser struct {
	UserID int `json:"name,required"`
}

func (h *authHandler) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload loginUser
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	err := h.userService.Login(ctx, payload.UserID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	// add cookie here
	respondJSON(w, http.StatusOK, nil)
}

func (h *authHandler) logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value(UserKey).(*core.User)
	if user == nil {
		respondError(w, http.StatusBadRequest,
			errors.New("logout: user already logged out"))
	}

	err := h.userService.Logoout(ctx, user.ID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	// delete cookie here
	respondJSON(w, http.StatusOK, nil)
}
