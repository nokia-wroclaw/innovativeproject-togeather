package server

import "net/http"

type pingHandler struct {
}

func (h *pingHandler) ping(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	respondJSON(w, http.StatusOK, "pong")
}