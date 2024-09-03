package main

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"

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

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("world 2"))
		if err != nil {
			return
		}
	})

	fmt.Println("listening on 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return
	}
}
