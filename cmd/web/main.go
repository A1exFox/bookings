package main

import (
	"fmt"
	"net/http"

	"github.com/a1exfox/go-course/pkg/handlers"
)

const portNumber = "localhost:8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
