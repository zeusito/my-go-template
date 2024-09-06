package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zeusito/my-go-template/pkg/router"
)

type HealthController struct{}

func NewHealthController(mux *chi.Mux) *HealthController {
	c := &HealthController{}

	mux.Get("/health/readiness", c.handleReadiness)
	mux.Get("/health/liveness", c.handleLiveness)

	return c
}

func (c *HealthController) handleReadiness(w http.ResponseWriter, req *http.Request) {
	router.RenderJSON(req.Context(), w, http.StatusOK, map[string]string{"status": "ok"})
}

func (c *HealthController) handleLiveness(w http.ResponseWriter, req *http.Request) {
	router.RenderJSON(req.Context(), w, http.StatusOK, map[string]string{"status": "ok"})
}
