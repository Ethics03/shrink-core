package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there!")
}

func main() {
    http.HandleFunc("/", handler)
		log.Printf("Listening to port: 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
		}

