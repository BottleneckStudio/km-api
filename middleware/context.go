package middleware

import (
	"net"
	"net/http"
	"os"
	"time"

	"context"

	"github.com/BottleneckStudio/km-api/services/post"
	"github.com/BottleneckStudio/km-api/services/user"

	// autoload values from .env file to ENV
	_ "github.com/joho/godotenv/autoload"
)

const PostServiceKey = "postService"
const UserServiceKey = "userService"
const ClientKey = "client"

var (
	dynamoTablePosts = os.Getenv("DYNAMO_TABLE_POSTS")
	dynamoEndpoint   = os.Getenv("DYNAMO_ENDPOINT")
	cognitoID        = os.Getenv("COGNITO_CLIENT_ID")
	cognitoSecret    = os.Getenv("COGNITO_CLIENT_SECRET")
)

// ClientContext middleware will inject the custom client to the context object
func ClientContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		}

		// support timeout and net transport.
		c := &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		}

		ctx := context.WithValue(r.Context(), ClientKey, c) //nolint

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// PostContext middleware will inject the post service to the context object
func PostContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := post.New(dynamoTablePosts, dynamoEndpoint, nil)

		ctx := context.WithValue(r.Context(), PostServiceKey, p) //nolint

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// UserContext middleware will inject the user service to the context object
func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := user.New(cognitoID, cognitoSecret)

		ctx := context.WithValue(r.Context(), UserServiceKey, u) //nolint

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
