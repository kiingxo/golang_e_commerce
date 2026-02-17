package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Application struct {
	Config Config
}

// mount
func (app *Application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return r

}

type Config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
