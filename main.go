package main

import (
	"github.com/caarlos0/env/v11"
	"log"
	"net/http"
)

type Config struct {
}

func main() {
	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			return
		}
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		return
	}
}
