package logging

import (
	"github.com/TnLCommunity/corndogs/server/config"
	"github.com/rs/zerolog/pkgerrors"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// InitLogging sets the TimestampFunc to current UTC time, and TimeFieldFormat to RFC339Nano
func InitLogging() {
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}
	zerolog.TimeFieldFormat = time.RFC3339Nano
	setLogLevel()
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func setLogLevel() {
	var level zerolog.Level
	switch strings.ToLower(config.LogLevel) {
	case "panic":
		level = zerolog.PanicLevel
	case "fatal":
		level = zerolog.FatalLevel
	case "error":
		level = zerolog.ErrorLevel
	case "warn":
		level = zerolog.WarnLevel
	case "info":
		level = zerolog.InfoLevel
	case "debug":
		level = zerolog.DebugLevel
	case "trace":
		level = zerolog.TraceLevel
	default:
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
}
