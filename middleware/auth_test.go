package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lestrrat-go/jwx/jwk"
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
	t.Run("Nil Keyset", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://randomsite.com", nil)

		AuthCheck(nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//something
		})).ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusInternalServerError)
	})

	t.Run("empty authorization header", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://randomsite.com", nil)

		AuthCheck(&jwk.Set{})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusUnauthorized)
	})
}
