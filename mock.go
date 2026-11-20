package logger

import "context"

// mock to logger.
type mock struct{}

func (m mock) Error(_ ...any) {
}

func (m mock) Errorf(_ string, _ ...any) {
}

func (m mock) Debug(_ ...any) {
}

func (m mock) Debugf(_ string, _ ...any) {
}

func (m mock) Info(_ ...any) {
}

func (m mock) Infof(_ string, _ ...any) {
}

func (m mock) Warn(_ ...any) {
}

func (m mock) Warnf(_ string, _ ...any) {
}

func (m mock) With(_ string, _ any) Logger {
	return m
}

func (m mock) WithError(_ error) Logger {
	return m
}

func (m mock) WithCorrelation(_ context.Context) Logger {
	return m
}

// Mock returns a new logger mock.
func Mock() Logger {
	return mock{}
}
