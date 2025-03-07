package log

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/response"
)

// Ctx brings together zerolog and gin.Context
func Ctx(ctx context.Context) *zerolog.Logger {
	logger := log.With()
	if requestID, ok := ctx.Value(response.RequestIDKey).(string); ok {
		logger = logger.Str(response.RequestIDKey, requestID)
	}
	if traceID, ok := ctx.Value(response.TraceIDKey).(string); ok {
		logger = logger.Str(response.TraceIDKey, traceID)
	}
	l := logger.Logger()
	return &l
}

func Info() *zerolog.Event {
	return log.Info()
}

func Error() *zerolog.Event {
	return log.Error()
}

func Warn() *zerolog.Event {
	return log.Warn()
}

func Debug() *zerolog.Event {
	return log.Debug()
}

func Trace() *zerolog.Event {
	return log.Trace()
}

func Panic() *zerolog.Event {
	return log.Panic()
}
