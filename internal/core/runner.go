package core

import (
	"context"
	"fmt"
	"sync"

	"github.com/salapati95/gometer/internal/config"
	"github.com/salapati95/gometer/internal/protocol"
	"github.com/salapati95/gometer/internal/zlog"
)

type Runner struct {
	cfg    *config.Config
	logger *zlog.Logger
}

func NewRunner(cfg *config.Config, logger *zlog.Logger) *Runner {
	return &Runner{cfg: cfg, logger: logger}
}

func (r *Runner) Start(ctx context.Context) error {
	r.logger.Info().
		Str("target", r.cfg.TargetHost).
		Int("port", r.cfg.TargetPort).
		Dur("duration", r.cfg.Duration).
		Int("connections", r.cfg.Connections).
		Msg("Starting load test")

	addr := fmt.Sprintf("%s:%d", r.cfg.TargetHost, r.cfg.TargetPort)
	client := &protocol.Client{
		Addr:     addr,
		Payload:  r.cfg.PayloadContent,
		Interval: r.cfg.Interval,
		Logger:   r.logger,
	}

	ctx, cancel := context.WithTimeout(ctx, r.cfg.Duration)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(r.cfg.Connections)

	for i := 0; i < r.cfg.Connections; i++ {
		go func(id int) {
			defer wg.Done()
			r.logger.Debug().Int("worker", id).Msg("Worker started")
			if err := client.Run(ctx); err != nil {
				r.logger.Warn().Err(err).Int("worker", id).Msg("Client run failed")
			}
		}(i)
	}

	wg.Wait()
	r.logger.Info().Msg("All workers completed")
	return nil
}
