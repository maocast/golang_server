package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func introduce(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Path: %s, MethodType: %s", r.URL.Path, r.Method)
    if r.Method == "POST" {
    	decoder := json.NewDecoder(r.Body)

		var p Person
	    err := decoder.Decode(&p)
	    if err != nil {
	    	fmt.Println(err)
	    }

	   	fmt.Println(p.FirstName)
	    fmt.Println(p.LastName)
    } else {
    	fmt.Fprintf(w, "Please use POST method", r.Method)
    }
}

type Person struct {
    FirstName, LastName string
}

func main() {
	http.HandleFunc("/introduce", introduce)
    http.ListenAndServe(":8080", nil)
}