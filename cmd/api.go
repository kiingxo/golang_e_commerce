package main

type Application struct {
	Config Config
}

// mount
func (app *Application) mount() {
	//mount handlers
}

type Config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
