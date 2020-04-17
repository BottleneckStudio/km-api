package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

// AuthCheck middleware will check validity of Token passed in request
func AuthCheck(keySet *jwk.Set) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if keySet == nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			tokenString := tokenFromHeader(r)
			if tokenString == "" {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				kid, ok := token.Header["kid"].(string)
				if !ok {
					return nil, errors.New("kid header not found")
				}
				keys := keySet.LookupKeyID(kid)
				if len(keys) == 0 {
					return nil, fmt.Errorf("key %v not found", kid)
				}
				return keys[0].Materialize()
			})

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
