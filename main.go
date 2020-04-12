package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BottleneckStudio/km-api/handler"
	mw "github.com/BottleneckStudio/km-api/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
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
			r.Use(mw.ClientContext)
			r.Use(mw.PostContext)
			r.Post("/", handler.CreatePost)
			r.Get("/", handler.GetPosts)
			r.Get("/{id}", handler.GetPost)
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
