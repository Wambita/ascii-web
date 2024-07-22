package main

import (
	"fmt"
	"net/http"
	"text/template"

	functions "server/functions"
)

type ascii struct {
	AsciiArt string
	Error    string
}

func home(w http.ResponseWriter, r *http.Request) {
	home := template.Must(template.ParseFiles("./template/index.html"))
	home.Execute(w, nil)
}

func ErrorPage(w http.ResponseWriter, statusCode int, message string) {
	tmpl := template.Must(template.ParseFiles("template/error.html"))
	w.WriteHeader(statusCode)
	data := ascii{Error: message}
	tmpl.Execute(w, data)
}

func art(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorPage(w, http.StatusBadRequest, "400 Bad Request")
		return
	}

	r.ParseForm()
	text := r.FormValue("input")
	banner := r.FormValue("bannerfile")
	fmt.Println((text))
	ascii_Art, err := functions.Input(text, banner)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	fmt.Println(ascii_Art)
	data := ascii{AsciiArt: ascii_Art}

	home := template.Must(template.ParseFiles("template/index.html"))
	err2 := home.Execute(w, data)
	if err2 != nil {
		fmt.Println(err2)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii-art", art)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is starting at http://localhost:8082")
	http.ListenAndServe(":8082", nil)
}
