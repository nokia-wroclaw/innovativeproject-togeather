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

func (h *restaurantHandler) exists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	restaurantID, err := strconv.Atoi(chi.URLParam(r, "restaurantID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	exists, err := h.restaurantService.Exists(ctx, restaurantID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, exists)
}

func (h *restaurantHandler) listRestaurants(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	restaurantsList, err := h.restaurantService.ListRestaurants(ctx)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, restaurantsList)
}

func (h *restaurantHandler) restaurantMenu(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	restaurantID, err := strconv.Atoi(chi.URLParam(r, "restaurantID"))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	restaurantMenu, err := h.restaurantService.RestaurantMenu(ctx, restaurantID)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	respondJSON(w, http.StatusOK, restaurantMenu)
}
