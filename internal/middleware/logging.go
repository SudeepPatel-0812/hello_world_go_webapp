package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type responseLogger struct {
	http.ResponseWriter
	status int
}

func (r *responseLogger) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Wrap the ResponseWriter with the custom responseLogger
		logger := &responseLogger{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(logger, r)

		duration := time.Since(startTime)
		fmt.Printf("%s %s - %d - %d ms\n", r.Method, r.URL.Path, logger.status, duration.Milliseconds())
	})
}
