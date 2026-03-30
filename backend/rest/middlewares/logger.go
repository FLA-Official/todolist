package middlewares

import (
	"log/slog"
	"net/http"
	"time"
	"todolist/utils"
)

func Logger(logger *slog.Logger) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			rw := &ResponseWriter{
				ResponseWriter: w,
				StatusCode:     http.StatusOK,
			}

			// Get Request ID
			reqID, _ := r.Context().Value(RequestIDKey).(string)

			reqLogger := logger.With(
				"request_id", reqID,
				"method", r.Method,
				"path", r.URL.Path,
			)

			// Attach user_id if available
			if payload, ok := r.Context().Value("user").(utils.Payload); ok {
				reqLogger = reqLogger.With("user_id", payload.ID)
			}

			// Inject logger into context BEFORE handler
			ctx := utils.WithLogger(r.Context(), reqLogger)

			// Call handler ONLY ONCE
			next.ServeHTTP(rw, r.WithContext(ctx))

			duration := time.Since(start)

			// Final log

			if rw.StatusCode >= 400 {
				reqLogger.Error(
					rw.errMsg, // 🔥 actual error message
					"status", rw.StatusCode,
					"duration", duration.String(),
					"ip", r.RemoteAddr,
				)
			} else {
				reqLogger.Info(
					"request completed",
					"status", rw.StatusCode,
					"duration", duration.String(),
					"ip", r.RemoteAddr,
				)
			}
		})
	}
}
