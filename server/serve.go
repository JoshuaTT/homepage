package server

import (
	// "fmt"
	"html/template"
	"net/http"
)

func handleHome(w http.ResponseWriter, req *http.Request) {
	templ := template.Must(template.ParseGlob("static/*.html"))

	templ.ExecuteTemplate(w, "index.html", nil)	
}

func Serve(addr string) {
	http.HandleFunc("/", handleHome)

	http.ListenAndServe(addr, nil)
}
