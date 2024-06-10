package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func AuthecationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Missing authentication token", http.StatusBadRequest)
			return
		}

		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authentication token", http.StatusBadRequest)
			return
		}

		token = tokenParts[1]

		claims, err := VerifyToken(token)
		if err != nil {
			http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
			return
		}

		userID := claims["user_id"].(float64)
		r.Header.Set("user_id", fmt.Sprintf("%d", int(userID)))

		next.ServeHTTP(w, r)
	})
}
