package main

import (
	"context"
	"log"

	"github.com/inRiza/popout/internal/config"
	authhandler "github.com/inRiza/popout/internal/handler/auth"
	"github.com/inRiza/popout/internal/pkg/supabase"
	authrepo "github.com/inRiza/popout/internal/repository/user"
	"github.com/inRiza/popout/internal/router"
	authservice "github.com/inRiza/popout/internal/service/auth"
	"github.com/inRiza/popout/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("[MAIN] failed to load config: %w", err)
	}

	pool, err := supabase.NewClient(context.Background(), cfg.DB.URL)
	if err != nil {
		log.Fatalf("[MAIN] failed to create supabase client: %w", err)
	}
	defer pool.Close()

	// wiring layers
	repo := authrepo.New(pool)
	svc := authservice.New(cfg, repo)
	handler := authhandler.New(cfg, svc)

	r := router.New(router.Deps{
		Config: cfg,
		AuthHandler: handler,
		AuthService: svc,
	})

	srv := server.New(cfg, r)
	if err := srv.Start(); err != nil {
		log.Fatalf("[MAIN] failed to start server: %w", err)
	}
}