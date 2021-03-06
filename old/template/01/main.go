package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {

	//テンプレートの解析
	t, _ := template.ParseFiles("tmpl.html")

	//レンダリング
	t.Execute(w, "Hello, World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
