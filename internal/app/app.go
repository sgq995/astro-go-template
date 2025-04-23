package app

import (
	"io/fs"
	"net/http"
)

type App struct {
	*http.ServeMux

	fs fs.FS
}

func New() *App {
	mux := http.NewServeMux()

	return &App{ServeMux: mux}
}

func (app *App) Init() error {
	return nil
}
