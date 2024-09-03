package main

import "net/http"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		write, err := w.Write([]byte("pong"))
		if err != nil {
			return
		}
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		return
	}
}
