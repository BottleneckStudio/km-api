package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/BottleneckStudio/km-api/services/user"
)

const (
	TokenKey = "token"
	UserKey  = "user"
)

// Token will check if there is an Authorization header and add it to the context
func Token(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := tokenFromHeader(r)
		if token != "" {
			ctx := context.WithValue(r.Context(), TokenKey, token) //nolint
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

// SessionUser middleware checks for the `user` session and if it exists
// it will try to fetch the user details from Cognito service and attach
// them to the context object.
func SessionUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			next.ServeHTTP(w, r)
		}()

		token := r.Context().Value(TokenKey)
		if token == nil {
			return
		}

		us := r.Context().Value(UserServiceKey).(interface {
			AccountDetails(token string) *user.User
		})

		user := us.AccountDetails(token.(string))
		if user == nil {
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, user) // nolint

		r = r.WithContext(ctx)
	})
}

func tokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
