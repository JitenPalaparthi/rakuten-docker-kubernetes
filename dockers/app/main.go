package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	println("Started and listening on 8080 port")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	http.ListenAndServe(":"+port, nil)
}
