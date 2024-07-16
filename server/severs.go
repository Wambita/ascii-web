package main

import (
	// "fmt"
	"handler/ascii-art/functions"
	"net/http"
	"text/template"
)

type PageData struct {
	AsciiArt string
}

func home(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	if r.Method == http.MethodPost {
		r.ParseForm()
		inputText := r.FormValue("inputText")
		asciiArt := functions.Input(inputText)
		data.AsciiArt = asciiArt
	}

	home := template.Must(template.ParseFiles("../template/index.html"))
	home.Execute(w, nil)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
