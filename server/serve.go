package server

import (
	"fmt"
	// "html/template"
	"net/http"
)

func handleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome home!")
}

func Serve(addr string) {
	http.HandleFunc("/", handleHome)

	http.ListenAndServe(addr, nil)
}
