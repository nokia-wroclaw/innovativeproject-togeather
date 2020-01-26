package server

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	user, err := h.userService.Create(ctx, payload.Name)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

type loginUser struct {
	UserID int `json:"id,required"`
}

func (h *authHandler) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload loginUser
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	u, err := h.userService.Login(ctx, payload.UserID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	id := strconv.Itoa(payload.UserID)

	c := &http.Cookie {
		Name:       CookieUserIDKey,
		Value:      id,
		Path: "/",
		MaxAge:     24*60*60,
		HttpOnly:   true,
	}
	http.SetCookie(w, c)
	respondJSON(w, http.StatusOK, u)
}

func (h *authHandler) logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "user-id=deleted; Path=/; Max-Age=0")
	respondJSON(w, http.StatusOK, nil)
}
