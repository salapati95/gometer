package config

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/salapati95/gometer/internal/zlog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	TargetHost     string        `yaml:"target_host"`
	TargetPort     int           `yaml:"target_port"`
	Duration       time.Duration `yaml:"duration"`
	Connections    int           `yaml:"connections"`
	PayloadPath    string        `yaml:"payload_path"`
	PayloadContent []byte        `yaml:"-"`
	Interval       time.Duration `yaml:"interval"`
}

// Load returns a default config for now. Later we can read from flags or env vars.
func Load(logger *zlog.Logger) (*Config, error) {
	configPath := os.Getenv("GOMETER_CONFIG")
	if configPath == "" {
		configPath = "config.yaml"
	}

	logger.Info().Str("path", configPath).Msg("Loading config")
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %w", configPath, err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		logger.Error().Err(err).Str("path", configPath).Msg("Failed to decode config")
		return nil, fmt.Errorf("could not decode %s: %w", configPath, err)
	}

	if cfg.PayloadPath == "" {
		logger.Error().Msg("Missing required field: payload_path")
		return nil, fmt.Errorf("payload_path is required in config")
	}

	payload, err := os.ReadFile(cfg.PayloadPath)
	if err != nil {
		logger.Error().Err(err).Str("payload_path", cfg.PayloadPath).Msg("Failed to read payload file")
		return nil, fmt.Errorf("could not read payload file: %w", err)
	}

	decoded, err := decodeHexPayload(payload)
	if err != nil {
		logger.Error().Err(err).Str("payload_path", cfg.PayloadPath).Msg("Invalid hex payload")
		return nil, fmt.Errorf("invalid hex payload: %w", err)
	}
	cfg.PayloadContent = decoded

	return &cfg, nil
}

func decodeHexPayload(data []byte) ([]byte, error) {
	decoded := make([]byte, hex.DecodedLen(len(data)))
	n, err := hex.Decode(decoded, bytes.TrimSpace(data))
	if err != nil {
		return nil, err
	}
	return decoded[:n], nil
}
