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

			next.ServeHTTP(rw, r)

			duration := time.Since(start)

			reqID, _ := r.Context().Value(RequestIDKey).(string)

			payload, ok := r.Context().Value("user").(utils.Payload)
			if ok {
				logger.Info("request",
					"request_id", reqID,
					"user_id", payload.ID,
					"method", r.Method,
					"path", r.URL.Path,
					"status", rw.StatusCode,
					"duration", duration.String(),
					"ip", r.RemoteAddr,
				)
			}

		})
	}
}
