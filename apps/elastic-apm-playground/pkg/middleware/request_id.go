package middleware

import (
	"context"
	"github.com/gorilla/mux"
	"go.elastic.co/apm/v2"
	"net/http"
)

const (
	XRequestID   = "X-Request-ID"
	requestIDCtx = "request_id"
)

func RequestIDMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if requestID := r.Header.Get(XRequestID); requestID != "" {
				transaction := apm.TransactionFromContext(r.Context())
				transaction.Context.SetLabel(XRequestID, requestID)

				ctx = context.WithValue(ctx, requestIDCtx, requestID)

				w.Header().Add(XRequestID, requestID)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequestIDFromCtx(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(requestIDCtx).(string)
	return requestID, ok
}
