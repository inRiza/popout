package auth

import (
	"context"

	"github.com/inRiza/popout/internal/config"
	"github.com/inRiza/popout/internal/model"
)

type Service struct {
	cfg *config.Config
	repo interface{}
}

func New(cfg *config.Config, repo interface{}) *Service {
	return &Service{cfg: cfg, repo: repo}
}

// TODO: implement
func (s *Service) ValidateSession(_ context.Context, _ string) (*model.User, error) {
	return nil, nil
}
