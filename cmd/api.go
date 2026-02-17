package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	Config Config
}

// mount
func (app *Application) mount() http.Handler {
	return chi.NewRouter()

}

type Config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
