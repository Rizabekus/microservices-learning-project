package middlewares

import (
	"net/http"
	"time"

	"github.com/Rizabekus/microservices-learning-project/order/internal/infrastructure/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.Log.Info("request started",
			"method", r.Method,
			"path", r.URL.Path,
		)

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		logger.Log.Info("request finished",
			"method", r.Method,
			"path", r.URL.Path,
			"duration_ms", duration.Milliseconds(),
		)
	})
}
