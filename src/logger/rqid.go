package logger

import (
	"context"

	"github.com/google/uuid"
)

// NewRequestID returns new request ID as string.
func NewRequestID() string {
	return uuid.New().String()
}

// WithRqID returns a context which knows its request ID
func WithRqID(ctx context.Context, rqID string) context.Context {
	return context.WithValue(ctx, RqIDCtxKey, rqID)
}
