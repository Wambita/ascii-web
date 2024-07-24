package main

import (
	"fmt"
	"net/http"
	"os"

	handler "server/handlers"
)

// type customHandler struct {
// 	http.Handler
// }

// func (c *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/", "/ascii-art", "/about", "/instructions":
// 		c.Handler.ServeHTTP(w, r)
// 	case "/static/", "/assets/":
// 		c.Handler.ServeHTTP(w, r)
// 	default:
// 		handler.ErrorPageHandler(w, http.StatusNotFound, "404 not found")
// 	}
// }

func main() {
	if len(os.Args) != 1 {
		return
	}
	fs := http.FileServer(http.Dir("static"))
	assets := http.FileServer(http.Dir("assets"))
	// mux := http.NewServeMux()
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/ascii-art", handler.ArtHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/instructions", handler.InstructionsHandler)
	// custom := &customHandler{Handler: mux}
	fmt.Println("Server is starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
