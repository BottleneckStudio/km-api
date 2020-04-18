package middleware

import (
	"net/http"
	"strings"

	"github.com/BottleneckStudio/km-api/services/auth"
	"github.com/dgrijalva/jwt-go"
)

// AuthCheck middleware will check validity of Token passed in request
func AuthCheck(parser auth.Parser) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if parser == nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			tokenString := tokenFromHeader(r)
			if tokenString == "" {
				http.Error(w, http.StatusText(401), 401)
				return
			}

			token, err := parser.Parse(tokenString)

			if err != nil {
				http.Error(w, http.StatusText(401), 401)
				return
			}

			if token == nil || !token.Valid {
				http.Error(w, http.StatusText(401), 401)
				return
			}

			// validate claims
			if _, ok := token.Claims.(jwt.MapClaims); !ok {
				http.Error(w, http.StatusText(401), 401)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func tokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
