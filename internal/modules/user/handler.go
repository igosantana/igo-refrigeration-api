package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}

	w.Header().Set("COntent-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
