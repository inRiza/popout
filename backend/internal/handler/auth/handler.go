package auth

import (
	"net/http"

	"github.com/inRiza/popout/internal/config"
)

type Service interface {
	ValidateSession(ctx interface{Done() <- chan struct{}}, token string) (interface{}, error)
}

type Handler struct {
	cfg *config.Config
	service interface{}
}

func New(cfg *config.Config, service interface{}) *Handler {
	return &Handler{cfg: cfg, service: service}
}

func (h *Handler) GoogleRedirect(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) GoogleCallback(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {}