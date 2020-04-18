package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BottleneckStudio/km-api/handler"
	mw "github.com/BottleneckStudio/km-api/middleware"
	"github.com/BottleneckStudio/km-api/services/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lestrrat-go/jwx/jwk"
)

func main() {
	port := ":" + os.Getenv("PORT")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", hello)

	r.Route("/api/v1", func(r chi.Router) {

		r.Route("/posts", func(r chi.Router) {
			r.Use(mw.PostContext)

			r.Get("/", handler.GetPosts)
			r.Get("/{id}", handler.GetPost)

			r.Group(func(r chi.Router) {
				r.Use(mw.AuthCheck(auth.New(initializeKeySets())))
				r.Post("/", handler.CreatePost)
			})
		})
	})

	log.Fatal(http.ListenAndServe(port, r))
}

func hello(w http.ResponseWriter, r *http.Request) {
	stage := os.Getenv("UP_STAGE")
	if stage == "" {
		stage = "Local"
	}
	fmt.Fprintf(w, "Hello!! API serving from %s", stage)
}

func initializeKeySets() *jwk.Set {
	content, err := ioutil.ReadFile("jwks.json")
	if err != nil {
		log.Printf("failed to read file: %s", err)
		return nil
	}
	set, err := jwk.ParseBytes(content)
	if err != nil {
		log.Printf("failed to parse JWK: %s", err)
		return nil
	}
	return set
}
