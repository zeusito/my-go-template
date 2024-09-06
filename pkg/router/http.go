package router

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type HTTPRouter struct {
	Mux *chi.Mux
	srv *http.Server
}

func NewHTTPRouter() *HTTPRouter {
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request models (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(20 * time.Second))

	return &HTTPRouter{
		Mux: router,
	}
}

func (s *HTTPRouter) Start() {
	// Listening address
	listeningAddr := ":3000"

	log.Info().Msgf("Server listening on port %s", ":3000")

	// Customizing the server
	s.srv = &http.Server{
		Addr:         listeningAddr,
		Handler:      s.Mux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	// Start the server
	_ = s.srv.ListenAndServe()
}

func (s *HTTPRouter) Shutdown(ctx context.Context) {
	log.Info().Msg("Server shutting down...")
	_ = s.srv.Shutdown(ctx)
}
