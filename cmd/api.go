package main

type Application struct {
	Config Config
}

type Config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
}
