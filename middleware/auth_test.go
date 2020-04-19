package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BottleneckStudio/km-api/services/user"
	"github.com/stretchr/testify/assert"
)

func TestHelper(t *testing.T) {
	t.Run("Empty bearer token", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "http://randomsite.com", nil)

		r.Header.Add("Authorization", "")
		token := tokenFromHeader(r)

		assert.Equal(t, token, "")
	})

	t.Run("No authorization header", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "http://randomsite.com", nil)

		token := tokenFromHeader(r)

		assert.Equal(t, token, "")
	})

	t.Run("Bearer token included", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "http://randomsite.com", nil)

		r.Header.Add("Authorization", "Bearer someransadjhkahakjshdahdsa")
		token := tokenFromHeader(r)

		assert.NotEmpty(t, token)
	})
}

func TestAuthCheck(t *testing.T) {
	t.Run("No User", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://randomsite.com", nil)

		CheckAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		})).ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusUnauthorized)
	})

	t.Run("User in context", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://randomsite.com", nil)

		ctx := context.WithValue(r.Context(), "user", &user.User{}) //nolint
		r = r.WithContext(ctx)

		CheckAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		})).ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusOK)
	})
}
