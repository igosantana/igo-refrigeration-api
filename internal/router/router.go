package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/igosantana/igo-refrigeration-api/internal/modules/user"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	// User layers
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	r.Mount("/users", user.RegisterRoutes(userHandler))

	return r
}
