package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"segoqu.com/astro-go-template/internal/app"
)

func main() {
	app := app.New()
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: "localhost:4321"})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, pattern := app.Handler(r)
		if pattern == "" {
			proxy.ServeHTTP(w, r)
		} else {
			app.ServeHTTP(w, r)
		}
	})

	s := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Printf("Starting server at %s", s.Addr)
	s.ListenAndServe()
}
