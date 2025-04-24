package main

import (
	"embed"
	"net/http"

	"segoqu.com/astro-go-template/internal/app"
)

//go:embed *[^{.go}]
var fs embed.FS

func main() {
	app := app.New()

	err := app.Init()
	if err != nil {
		panic(err)
	}

	s := http.Server{
		Addr:    "localhost:8080",
		Handler: app,
	}

	s.ListenAndServe()
}
