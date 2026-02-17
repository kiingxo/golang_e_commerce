package main

func main() {
	cfg := Config{
		addr: ":8080",
		db:   dbConfig{},
	}
	api := Application{
		Config: cfg,
	}
}
