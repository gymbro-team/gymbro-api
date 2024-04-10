package auth

import (
	"gymbro-api/config"
	"net/http"
)

func AuthecationMiddleware(next http.Handler) http.Handler {
	cfg := config.LoadConfig()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != "Bearer "+cfg.Token {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
