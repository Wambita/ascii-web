package main

import (
	"fmt"
	"net/http"
	"os"

	handler "server/handlers"
)



// Runs the server and calls the various webpages.
func main() {
	if len(os.Args) != 1 {
		return
	}

	fs := http.FileServer(http.Dir("static"))
	assets := http.FileServer(http.Dir("assets"))


	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.HandleFunc("/ascii-art", handler.ArtHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/instructions", handler.InstructionsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/home":
			handler.HomeHandler(w, r)
		default:
			handler.ErrorPageHandler(w, http.StatusNotFound, "404 - Not Found")
		}
	})
	fmt.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func Serve(w http.ResponseWriter, r *http.Request) {
	imagePath := "./" + r.URL.Path
	_, err := os.Stat(imagePath)
	if err != nil {
		fmt.Fprintf(w, "Image not found")
	}
	
	content, _ := os.ReadFile(imagePath)
	fmt.Fprintf(w, "%s", string(content))
}

