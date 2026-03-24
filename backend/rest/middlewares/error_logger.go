package middlewares

import (
	"log/slog"
	"net/http"
)

func ErrorLogger(logger *slog.Logger) func(AppHandler) http.Handler {

	return func(next AppHandler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			err := next(w, r)

			if err == nil {
				return
			}

			reqID, _ := r.Context().Value(RequestIDKey).(string)

			logger.Error("handler error",
				"request_id", reqID,
				"method", r.Method,
				"path", r.URL.Path,
				"error", err,
			)

			http.Error(w, err.Error(), http.StatusInternalServerError)

		})
	}
}
