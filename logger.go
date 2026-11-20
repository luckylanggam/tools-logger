package logger

import (
	"context"

	"go.uber.org/zap"
)

type key string

// CorrelationIDKey is the key to access the correlation id in the context.
const CorrelationIDKey key = "correlation_id"

const (
	errorKey   string = "error"
	serviceKey string = "service"
)

// Logger to print information to standard output.
type Logger interface {
	Error(args ...any)
	Errorf(template string, args ...any)
	Debug(args ...any)
	Debugf(template string, args ...any)
	Info(args ...any)
	Infof(template string, args ...any)
	Warn(args ...any)
	Warnf(template string, args ...any)
	With(key string, value any) Logger
	WithError(err error) Logger
	WithCorrelation(ctx context.Context) Logger
}

// logger is an adapted zap logger.
type logger struct {
	*zap.SugaredLogger
	config *zap.Config
}

// With adds fields to the log.
func (l *logger) With(key string, value any) Logger {
	return &logger{
		config:        l.config,
		SugaredLogger: l.SugaredLogger.With(key, value),
	}
}

// WithError adds an error to the log context.
func (l *logger) WithError(err error) Logger {
	return l.With(errorKey, err)
}

// WithCorrelation adds the correlation id to the log if it is found in the context.
func (l *logger) WithCorrelation(ctx context.Context) Logger {
	// add our correlation id if present.
	cid := ctx.Value(CorrelationIDKey)
	if cid == nil {
		return l
	}

	return l.With(string(CorrelationIDKey), cid)
}

// New creates a new logger.
func New(serviceName string, debug bool) Logger {
	if !debug {
		config := zap.NewProductionConfig()
		l, _ := config.Build()
		return &logger{l.Sugar().With(zap.String(serviceKey, serviceName)), &config}
	}

	config := zap.NewDevelopmentConfig()
	l, _ := config.Build()
	return &logger{l.Sugar().With(zap.String(serviceKey, serviceName)), &config}
}
