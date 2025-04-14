package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/igosantana/igo-refrigeration-api/internal/modules/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(db *pgxpool.Pool) http.Handler {
	r := chi.NewRouter()

	// User layers
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	user.RegisterRoutes(r.Route("/users", nil), userHandler)

	return r
}
