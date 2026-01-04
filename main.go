package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprint(w, "Hello world")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("bytes written: %d\n", n)
	})

	http.ListenAndServe(":8080", nil)
}
