package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the ultra-fast Go API!")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC1123)
		fmt.Fprintf(w, "Current time: %s", currentTime)
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("msg")
		if message == "" {
			fmt.Fprintf(w, "Please provide a message to echo using the 'msg' query parameter.")
		} else {
			fmt.Fprintf(w, "Echo: %s", message)
		}
	})

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}

