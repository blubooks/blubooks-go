package router

import (
	"blubooks/server/app"

	//"blubooks/server/app/middleware"

	"github.com/go-chi/chi"
	//"github.com/go-chi/cors"
)

func New(a *app.App) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/healthz", app.HandleHealth)
	r.Get("/healthz2", app.HandleHealth)
	//r.Use(middleware.Logger("", nil))

	return r
}
