package infrastructure

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger

func SetupLogger(config Config) *zerolog.Logger {
	var writers []io.Writer
	dailyWriter := &DailyWriter{
		Dir:        config.LogDirectory,
		Compress:   true,
		ReserveDay: 7,
	}
	writers = append(writers, dailyWriter)

	if config.RunningEnv == "development" {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	mw := io.MultiWriter(writers...)

	logger := zerolog.New(mw).With().Timestamp().Logger()
	return &logger
}
