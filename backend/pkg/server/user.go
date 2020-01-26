package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type userHandler struct {
	userService core.UserService
}

func (h *userHandler) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	usersList, err := h.userService.List(ctx)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, usersList)
}

func (h *userHandler) get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := r.Cookie("user-id")
	if err != nil {
		respondError(w, http.StatusBadRequest,
			errors.New("user not authorized"))
		return
	}

	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.userService.Get(ctx, userID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, user)
}
