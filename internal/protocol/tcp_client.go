package protocol

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/salapati95/gometer/internal/zlog"
)

type Client struct {
	Addr     string
	Payload  []byte
	Interval time.Duration
	Logger   *zlog.Logger
}

func (c *Client) Run(ctx context.Context) error {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		c.Logger.Error().Err(err).Str("addr", c.Addr).Msg("Failed to connect")
		return err
	}
	defer conn.Close()

	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			c.Logger.Debug().Msg("Client context cancelled")
			return nil
		case <-ticker.C:
			_, err := conn.Write(c.Payload)
			if err != nil {
				c.Logger.Error().Err(err).Msg("Write failed")
				return fmt.Errorf("write error: %w", err)
			}
		}
	}
}
