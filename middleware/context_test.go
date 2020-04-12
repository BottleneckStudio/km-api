package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextMw(t *testing.T) {
	t.Run("should attach PostService to context", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)

		PostContext(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ps := ctx.Value(PostServiceKey)
			assert.NotNil(t, ps)
		})).ServeHTTP(w, r)

	})

	t.Run("should attach Client to context", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)

		ClientContext(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			client := ctx.Value(ClientKey)
			assert.NotNil(t, client)
		})).ServeHTTP(w, r)

	})
}
