package main

import (
    "fmt"
    "net/http"
)

func introduce(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Path: %s, MethodType: %s", r.URL.Path, r.Method)
}

func main() {
	http.HandleFunc("/introduce", introduce)
    http.ListenAndServe(":8080", nil)
}