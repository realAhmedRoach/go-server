package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"sces/api/modules"
	"sces/mgmt"
)

func Routes(app *mgmt.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	router.Mount("/api/v1", modules.OrderRoutes(app))

	return router
}
