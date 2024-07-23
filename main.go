package main

import (
	"fmt"
	"net/http"
	"os"
	handler "server/handlers"
)

func main() {
	if len(os.Args) != 1 {
		return
	}
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/ascii-art", handler.Art)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/assets/", http.StripPrefix("/assets/", nil))
	http.HandleFunc("/about", handler.About)
	http.HandleFunc("/instructions", handler.Instructions)
	fmt.Println("Server is starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
