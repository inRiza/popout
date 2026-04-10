package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/inRiza/popout/internal/config"
)

type Server struct {
	http *http.Server
}

func New(cfg *config.Config, handler http.Handler) *Server {
	return &Server {
		http: &http.Server {
			Addr: fmt.Sprintf(":%s", cfg.App.Port),
			Handler: handler,
			ReadTimeout: 15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout: 30 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)

	go func() {
		log.Printf("server is starting on %s", s.http.Addr)
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case err := <- errCh:
		return err
	case <- ctx.Done():
		log.Println("server is shutting down")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.http.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("[SERVER] failed to shutdown: %w", err)
	}

	log.Println("server shutdown completed")
	return nil
}