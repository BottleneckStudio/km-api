package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BottleneckStudio/km-api/services/post"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPosts(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		p := new(postMock)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)

		ctx := context.WithValue(r.Context(), "postService", p) //nolint
		r = r.WithContext(ctx)
		p.On("GetPosts").Return([]*post.Post{})

		GetPosts(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		p.AssertExpectations(t)
	})
}

func TestGetPost(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		p := new(postMock)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)

		ctx := context.WithValue(r.Context(), "postService", p) //nolint
		r = r.WithContext(ctx)
		p.On("GetPost", mock.MatchedBy(func(val string) bool {
			return true
		})).Return(nil)

		GetPost(w, r)

		assert.Equal(t, http.StatusNotFound, w.Code)
		p.AssertExpectations(t)
	})

	t.Run("ok", func(t *testing.T) {
		p := new(postMock)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)

		ctx := context.WithValue(r.Context(), "postService", p) //nolint
		r = r.WithContext(ctx)
		p.On("GetPost", mock.MatchedBy(func(val string) bool {
			return true
		})).Return(&post.Post{
			Title:   "test",
			Content: "test",
			ID:      "test",
		})

		GetPost(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		p.AssertExpectations(t)
	})
}

func TestCreatePost(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		p := new(postMock)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)

		ctx := context.WithValue(r.Context(), "postService", p) //nolint
		r = r.WithContext(ctx)
		CreatePost(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("ok", func(t *testing.T) {
		p := new(postMock)

		m := map[string]interface{}{
			"title":   "test",
			"content": "icles",
		}
		body, _ := json.Marshal(m)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/v1/posts", bytes.NewReader(body))

		ctx := context.WithValue(r.Context(), "postService", p) //nolint
		r = r.WithContext(ctx)

		p.On("CreatePost", mock.MatchedBy(func(vals map[string]interface{}) bool {
			return true
		})).Return(&post.Post{
			Title:   "test",
			Content: "test",
			ID:      "test",
		})
		CreatePost(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
		p.AssertExpectations(t)
	})

	t.Run("ok but nil response", func(t *testing.T) {
		p := new(postMock)

		m := map[string]interface{}{
			"title":   "test",
			"content": "icles",
		}
		body, _ := json.Marshal(m)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/v1/posts", bytes.NewReader(body))

		ctx := context.WithValue(r.Context(), "postService", p) //nolint
		r = r.WithContext(ctx)

		p.On("CreatePost", mock.MatchedBy(func(vals map[string]interface{}) bool {
			return true
		})).Return(nil)
		CreatePost(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		p.AssertExpectations(t)
	})
}
