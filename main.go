package main

import (
	"fmt"
	"net/http"
	handler "server/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/ascii-art", handler.Art)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is starting at http://localhost:8082")
	http.ListenAndServe(":8082", nil)
}
