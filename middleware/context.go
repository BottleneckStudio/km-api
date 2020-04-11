package middleware

import (
	"net"
	"net/http"
	"os"
	"time"

	"context"

	"github.com/BottleneckStudio/km-api/services/post"
)

var (
	cognitoID        = os.Getenv("COGNITO_CLIENT_ID")
	cognitoSecret    = os.Getenv("COGNITO_CLIENT_SECRET")
	dynamoTablePosts = os.Getenv("DYNAMO_TABLE_POSTS")
	dynamoTableLikes = os.Getenv("DYNAMO_TABLE_LIKES")
	dynamoEndpoint   = os.Getenv("DYNAMO_ENDPOINT")
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

		ctx := context.WithValue(r.Context(), "client", c)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// PostContext middleware will inject the post service to the context object
func PostContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := post.New(dynamoTablePosts, dynamoEndpoint, nil)

		ctx := context.WithValue(r.Context(), "postService", p)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
