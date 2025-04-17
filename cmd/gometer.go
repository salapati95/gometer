package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/salapati95/gometer/internal/config"
	"github.com/salapati95/gometer/internal/core"
	"github.com/salapati95/gometer/internal/zlog"
)

func main() {
	log := zlog.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Info().Msg("Shutting down...")
		cancel()
	}()

	cfg, err := config.Load(log)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	runner := core.NewRunner(cfg, log)
	if err := runner.Start(ctx); err != nil {
		log.Fatal().Err(err).Msg("Error running load test")
	}

	time.Sleep(1 * time.Second)
	log.Info().Msg("[gometer] Exited cleanly.")
}
