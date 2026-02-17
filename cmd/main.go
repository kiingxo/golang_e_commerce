package main

import "os"

func main() {
	cfg := Config{
		addr: ":8080",
		db: dbConfig{
			dsn: os.Getenv(""),
		},
	}
	api := Application{
		Config: cfg,
	}
}
