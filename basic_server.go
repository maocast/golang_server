package main

import (
	"encoding/json"
	"io"
	"net/http"
)

//Structures used for encoding/decoding json
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
			//Manually writing error message in json format. Should encode data
			http.Error(w, "{\"Message\": \""+err.Message+"\"}", http.StatusBadRequest)
			return
		}

		//There are no problems, generate response
		g := Greeter{"Hi, " + p.FirstName + " " + p.LastName}
		json.NewEncoder(w).Encode(&g)

	} else {
		//Manually writing error message in json format. Should encode data
		http.Error(w, "{\"Message\": \"Method not allowed\"}", http.StatusMethodNotAllowed)
	}
}

func decodeBody(r io.Reader) (*Person, *ServerError) {
	var p Person
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&p)

	//Problem decoding data
	if err != nil {
		return nil, &ServerError{"Could not load json data"}
	}
	//App specific special cases
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
