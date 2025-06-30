package middlewares

import (
	"net/http"
	"strings"

	"github.com/FedotCompot/file-cacher/internal/config"
	"github.com/FedotCompot/file-cacher/internal/web/api"
	"github.com/golang-jwt/jwt/v5"
	"github.com/uptrace/bunrouter"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.Data.JWTSecret), nil
}

// Middleware for JWT authentication via the Authorization header.
func JWTAuthMiddleware() bunrouter.MiddlewareFunc {

	parser := jwt.NewParser()

	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return func(w http.ResponseWriter, req bunrouter.Request) error {
			authHeader := req.Header.Get("Authorization")
			if authHeader == "" {
				return api.RenderStatus(w, http.StatusUnauthorized)
			}

			// Check if the header starts with "Bearer "
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && strings.ToLower(parts[0]) == "bearer") {
				return api.RenderStatus(w, http.StatusUnauthorized)
			}

			tokenString := parts[1]

			// Parse and validate the token
			token, err := parser.Parse(tokenString, keyFunc)
			if err != nil {
				return api.RenderStatus(w, http.StatusUnauthorized)
			}

			if !token.Valid {
				return api.RenderStatus(w, http.StatusUnauthorized)
			}

			// Call the next handler in the chain
			return next(w, req)
		}
	}
}
