package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)

	data["hoge"] = "fuga"
	data["piyo"] = "moge"

	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, data)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
