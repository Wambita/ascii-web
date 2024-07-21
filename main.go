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

func art(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	r.ParseForm()
	text := r.FormValue("input")
	banner := r.FormValue("bannerfile")
	fmt.Println((text))
	ascii_Art, err := functions.Input(text, banner)
	if err != nil {
		http.Error(w, "Error generating ASCII art", http.StatusInternalServerError)
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

func downloadAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusBadRequest)
		return
	}
	r.ParseForm()
	text := r.FormValue("input")
	banner := r.FormValue("bannerfile")
	fmt.Println((text))
	ascii_Art, err := functions.Input(text, banner)
	if err != nil {
		http.Error(w, "Error generating ASCII art", http.StatusInternalServerError)
		return
	}

	filename := "ascii-art.txt"
	err = functions.Download(ascii_Art, filename)
	if err != nil {
		http.Error(w, "Error saving ASCII art to file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Disposition", "attachement;filename="+filename)
	http.ServeFile(w, r, filename)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii-art", art)
	http.HandleFunc("/download-ascii-art", downloadAscii)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is starting at http://localhost:8082")
	http.ListenAndServe(":8082", nil)
}
