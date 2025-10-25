package ride

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// 🌍 CORS middleware
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// 🧾 Logging middleware
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[HTTP] %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("[HTTP] Completed in %v", time.Since(start))
	})
}

// 🪪 Auth middleware (простая проверка токена)
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// пример: "Bearer secret-token"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[1] != "secret-token" {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
