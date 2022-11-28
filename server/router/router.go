package router

import (
	"blubooks/server/app"
	"blubooks/server/app/middleware"

	//"blubooks/server/app/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func New(a *app.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger("", nil))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		//AllowedOrigins: []string{"https://*", "http://*"},
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Get("/healthz", app.HandleHealth)
	r.Get("/healthz2", app.HandleHealth)
	//r.Use(middleware.Logger("", nil))
	r.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.ContentTypeJson)
			r.Post("/auth/login", a.Login)
			r.Post("/auth/refresh", a.RefreshLoginToken)
		})

		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth(a))
			r.Use(middleware.ContentTypeJson)
			r.Get("/clients/{id}/collections", a.GetCollections)

			r.Get("/clients", a.HandleListClients)
		})
	})
	r.Get("/*", a.HandleIndex)
	return r
}
