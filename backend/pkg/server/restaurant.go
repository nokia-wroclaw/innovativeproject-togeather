package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type restaurantHandler struct {
	restaurantService core.RestaurantService
}

func (h *restaurantHandler) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	restaurantsList, err := h.restaurantService.List(ctx)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, restaurantsList)
}

func (h *restaurantHandler) menu(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	restaurantID, err := strconv.Atoi(chi.URLParam(r, "restaurantID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	restaurantMenu, err := h.restaurantService.GetMenu(ctx, restaurantID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, restaurantMenu)
}

func (h *restaurantHandler) get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	restaurantID, err := strconv.Atoi(chi.URLParam(r, "restaurantID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	restaurant, err := h.restaurantService.Get(ctx, restaurantID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, restaurant)
}
