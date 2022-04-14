package main

import (
	"fmt"
	// "net/http"
	// "html/template"
	"github.com/joshuatt/homepage/server"
)

const port = 8080
const addr = "localhost:8080"

func main() {
	fmt.Println(port)
	server.Serve(addr)
}
