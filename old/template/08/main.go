package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World!")
}

func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t3.html", "t2.html")
	t.Execute(w, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)

	server.ListenAndServe()
}
