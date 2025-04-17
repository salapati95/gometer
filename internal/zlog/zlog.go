package zlog

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func New() *Logger {
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	return &Logger{Logger: &l}
}
