package server

import (
	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
	"net/http"
)

type userHandler struct {
	userService core.UserService
}

func (h *userHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	usersList, err := h.userService.ListUsers(ctx)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, usersList)
}
