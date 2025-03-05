package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statTime := time.Now()
		log.Printf("Started %s Methos , PATH %s ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Time taken for request %v", time.Since(statTime))

	})
}
