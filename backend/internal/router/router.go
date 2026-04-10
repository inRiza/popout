package router

import (
	"net/http"
	
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/inRiza/popout/internal/config"
	authhandler "github.com/inRiza/popout/internal/handler/auth"
	"github.com/inRiza/popout/internal/middleware"
)

// dependencies for the router struct to use
type Deps struct {
	Config 			*config.Config
	AuthHandler 	*authhandler.Handler
	AuthService 	middleware.SessionValidator
}

func New(deps Deps) http.Handler {
	r := chi.NewRouter() // create router
	r.Use(chimiddleware.RealIP) // get real ip
	r.Use(chimiddleware.RequestID) // get request id
	r.Use(chimiddleware.Recoverer) // recover from panics
	r.Use(middleware.CORS(deps.Config)) // cors

	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			// public
			r.Get("/google", deps.AuthHandler.GoogleRedirect)
			r.Get("/google/callback", deps.AuthHandler.GoogleCallback)

			// protected
			r.Group(func(r chi.Router) {
				r.Use(middleware.ValidateSession(deps.AuthService))
				r.Post("/logout", deps.AuthHandler.Logout)
				r.Get("/me", deps.AuthHandler.Me)
			})
		})
	})

	return r
}