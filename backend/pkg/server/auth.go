package server

import (
	"encoding/json"
	"errors"
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
	ctx := r.Context()

	_, ok := ctx.Value(UserKey).(*core.User)
	if !ok {
		respondError(w, http.StatusBadRequest, errors.New("unauthorized"))
		return
	}

	_, err := r.Cookie(CookieUserIDKey)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	//http.SetCookie(w, c)
	addCookie(w, CookieUserIDKey, "", -1, "/")
	respondJSON(w, http.StatusNoContent, nil)
}
