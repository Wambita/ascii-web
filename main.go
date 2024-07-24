package main

import (
	"net/http"
	"os"

	handler "server/handlers"
)

func main() {
	if len(os.Args) != 1 {
		return
	}

	fs := http.FileServer(http.Dir("static"))
	assets := http.FileServer(http.Dir("assets"))
	// mux := http.NewServeMux()
	http.HandleFunc("/home", handler.HomeHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.HandleFunc("/ascii-art", handler.ArtHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/instructions", handler.InstructionsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request path matches any of the existing handlers
		if r.URL.Path == "/" ||
		r.URL.Path == "/home" ||
			r.URL.Path == "/ascii-art" ||
			r.URL.Path == "/about" ||
			r.URL.Path == "/instructions" {
			http.NewServeMux().ServeHTTP(w, r)
		} else {
			handler.ErrorPageHandler(w, http.StatusNotFound, "404 - Not Found")
		}
	})

	// custom := &customHandler{Handler: mux}

	http.ListenAndServe(":8080", nil)
}
