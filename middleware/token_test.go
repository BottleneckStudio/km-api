package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BottleneckStudio/km-api/services/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestToken(t *testing.T) {
	t.Run("No token", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "https://somethingsomething.com", nil)

		Token(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Context().Value(TokenKey)
			assert.Nil(t, token)
		})).ServeHTTP(w, r)
	})

	t.Run("Token available", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "https://somethingsomething.com", nil)

		r.Header.Add("Authorization", "Bearer someransadjhkahakjshdahdsa")
		Token(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Context().Value(TokenKey)
			assert.NotNil(t, token)
		})).ServeHTTP(w, r)
	})
}

func TestSessionUser(t *testing.T) {
	t.Run("No token", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "https://asdadadad", nil)

		SessionUser(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(UserKey)
			assert.Nil(t, user)
		})).ServeHTTP(w, r)
	})

	t.Run("Nil User service", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "https://asdadadad", nil)

		r.Header.Add("Authorization", "Bearer someransadjhkahakjshdahdsa")

		us := new(userServiceMock)
		ctx := context.WithValue(r.Context(), UserServiceKey, us) // nolint
		r = r.WithContext(ctx)

		us.On("AccountDetails", mock.MatchedBy(func(tokenString string) bool {
			return true
		})).Return(nil)

		Token(SessionUser(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(UserKey)
			assert.Nil(t, user)
		}))).ServeHTTP(w, r)
	})

	t.Run("ok", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "https://asdadadad", nil)

		r.Header.Add("Authorization", "Bearer someransadjhkahakjshdahdsa")

		us := new(userServiceMock)
		ctx := context.WithValue(r.Context(), UserServiceKey, us) // nolint
		r = r.WithContext(ctx)

		us.On("AccountDetails", mock.MatchedBy(func(tokenString string) bool {
			return true
		})).Return(&user.User{})

		Token(SessionUser(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(UserKey)
			assert.NotNil(t, user)
		}))).ServeHTTP(w, r)
	})
}
