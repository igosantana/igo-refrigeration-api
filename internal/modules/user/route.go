package user

import "github.com/go-chi/chi/v5"

func RegisterRoutes(h *Handler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.GetAll)
	return r
}
