package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start).Milliseconds()
		str := fmt.Sprintf("%-7s %-20s %dms", r.Method, r.URL, elapsed)
		fmt.Println(str)
	})
}
