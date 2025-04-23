package app

import "net/http"

type App struct {
	*http.ServeMux
}

func New() *App {
	mux := http.NewServeMux()

	// TODO: register routes

	return &App{ServeMux: mux}
}
