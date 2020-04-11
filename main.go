package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	port := ":" + os.Getenv("PORT")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", hello)
	r.Get("/greet/{name}", greeter)

	log.Fatal(http.ListenAndServe(port, r))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func greeter(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	fmt.Fprintf(w, "Hi there %v!", name)
}
