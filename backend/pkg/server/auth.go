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

	err := h.userService.Login(ctx, payload.UserID)
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
	respondJSON(w, http.StatusOK, payload.UserID)
}

func (h *authHandler) logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value(UserKey).(*core.User)
	if user == nil {
		respondError(w, http.StatusBadRequest,
			errors.New("logout: user already logged out"))
	}

	c := &http.Cookie {
		Name:       CookieUserIDKey,
		Value:      "",
		Path: 		"/",
		MaxAge:     -1,
		HttpOnly:   true,
	}

	//err := h.userService.Logout(ctx, user.ID)
	//if err != nil {
	//	respondError(w, http.StatusBadRequest, err)
	//	return
	//}

	// delete cookie here
	http.SetCookie(w, c)
	respondJSON(w, http.StatusOK, nil)
}
