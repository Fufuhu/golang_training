package main

import (
	"html/template"
	"net/http"
)

func redProcess(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html", "red_hello.html")
	t.ExecuteTemplate(w, "layout", "")
}

func blueProcess(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html", "blue_hello.html")
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/red", redProcess)
	http.HandleFunc("/blue", blueProcess)

	server.ListenAndServe()
}
