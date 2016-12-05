package main

import (
    "net/http"
    "encoding/json"
    "io"
)

type Person struct {
    FirstName, LastName string
}

type Greeter struct {
    Body string
}

type ServerError struct {
	Message string
}

func introduce(w http.ResponseWriter, r *http.Request) {
    //Only allow for post methods
    if r.Method == "POST" {
    	p, err := decodeBody(r.Body)
    	//If there was problem decoding the body parameters
    	if err != nil {
	    	http.Error(w, "{\"Message\": \"" + err.Message + "\"}", http.StatusBadRequest)
	    	return
    	}

    	//There are no problems, generate response
		g := Greeter {"Hi, " + p.FirstName + " " + p.LastName}
		json.NewEncoder(w).Encode(&g)

    } else {
    	http.Error(w, "{\"Message\": \"Method not allowed\"}", http.StatusMethodNotAllowed)
    }
}

func decodeBody(r io.Reader) (*Person, *ServerError) {
    var p Person
    decoder := json.NewDecoder(r)
    err := decoder.Decode(&p)
    
    if err != nil {
    	return nil, &ServerError{"Could not load json data"}
    }
    if p.FirstName == "" {
    	return nil, &ServerError{"Person must have a first name"}
    }
    if p.LastName == "" {
    	return nil, &ServerError{"Person must have a last name"}
    }

    return &p, nil
}

func main() {
	http.HandleFunc("/introduce", introduce)
    http.ListenAndServe(":8080", nil)
}