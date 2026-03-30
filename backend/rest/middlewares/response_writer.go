package middlewares

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	errMsg     string
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.StatusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	// Capture error body (only for error responses)
	if rw.StatusCode >= 400 && rw.errMsg == "" {
		rw.errMsg = string(b)
	}
	return rw.ResponseWriter.Write(b)
}
