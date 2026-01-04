package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "This is home page, sum: %d", sum)
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(3, 3)
	fmt.Fprintf(w, "This is about page, sum: %d", sum)
}

func addValues(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)

	http.ListenAndServe(portNumber, nil)
}
