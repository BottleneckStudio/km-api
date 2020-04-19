package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	test "github.com/BottleneckStudio/km-api/testing"
	"github.com/dgrijalva/jwt-go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	t.Run("Nil Parser", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://randomsite.com", nil)

		AuthCheck(nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		})).ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusInternalServerError)
	})

	t.Run("empty authorization header", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://randomsite.com", nil)
		m := new(test.ParserMock)

		m.On("Parse", mock.MatchedBy(func(tokenString string) bool {
			return true
		})).Return(nil, errors.New("Empty AuthHeader"))

		AuthCheck(m)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusUnauthorized)
	})

	t.Run("Nil token", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://keepmotivatin", nil)

		r.Header.Add("Authorization", "Bearer eyyyzjasdqe.someransadjhkahakjshdahdsa")
		m := new(test.ParserMock)

		m.On("Parse", mock.MatchedBy(func(tokenString string) bool {
			return true
		})).Return(nil, errors.New("Some error"))

		AuthCheck(m)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, r)
		assert.Equal(t, w.Code, http.StatusUnauthorized)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://keepmotivatin", nil)

		r.Header.Add("Authorization", "Bearer eyyyzjasdqe.someransadjhkahakjshdahdsa")
		m := new(test.ParserMock)

		m.On("Parse", mock.MatchedBy(func(tokenString string) bool {
			return true
		})).Return(&jwt.Token{
			Valid: false,
		}, nil)

		AuthCheck(m)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusUnauthorized)
	})

	t.Run("Invalid Claims", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://keepmotivatin", nil)

		r.Header.Add("Authorization", "Bearer eyyyzjasdqe.someransadjhkahakjshdahdsa")
		m := new(test.ParserMock)

		m.On("Parse", mock.MatchedBy(func(tokenString string) bool {
			return true
		})).Return(&jwt.Token{
			Valid: true,
		}, nil)

		AuthCheck(m)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			assert.Equal(t, r, req)
		})).ServeHTTP(w, r)

		assert.Equal(t, w.Code, http.StatusUnauthorized)
	})

	t.Run("ok", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://keepmotivatin", nil)

		r.Header.Add("Authorization", "Bearer eyyyzjasdqe.someransadjhkahakjshdahdsa")
		m := new(test.ParserMock)

		m.On("Parse", mock.MatchedBy(func(tokenString string) bool {
			return true
		})).Return(&jwt.Token{
			Valid:  true,
			Claims: make(jwt.MapClaims),
		}, nil)

		AuthCheck(m)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		})).ServeHTTP(w, r)

		m.AssertExpectations(t)
	})

}
