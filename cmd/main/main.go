package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/antremis/server/cmd/email"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{}))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(30 * time.Second))

	r.Route("/email", email.Router)

	fmt.Println("Starting Server")
	http.ListenAndServe(":10000", r)
}
