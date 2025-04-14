package user

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, handler *Handler) {
	r.Get("/", handler.GetAll)
}
