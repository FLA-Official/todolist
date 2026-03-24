package middlewares

import (
	"log/slog"
	"net/http"
)

func Recovery(logger *slog.Logger) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			defer func() {

				if err := recover(); err != nil {

					reqID, _ := r.Context().Value(RequestIDKey).(string)

					logger.Error("panic recovered",
						"request_id", reqID,
						"error", err,
						"path", r.URL.Path,
					)

					http.Error(w, "internal server error", http.StatusInternalServerError)
				}

			}()

			next.ServeHTTP(w, r)
		})
	}
}
