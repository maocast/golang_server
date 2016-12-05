package main

import (
	"net/http"
	"testing"
	"bytes"
	"io/ioutil"
	"strings"
)

func TestMainGet(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/introduce")
	if err != nil { //Should fail
		t.Error("Was not able to communicate with server")
	} else if resp.StatusCode != 405 { //Should fail
		t.Error("Only post method should be allowed")
	}
}

func TestMainPostBadFormat(t *testing.T) {
	buf := bytes.NewBufferString("{\"titleasdf")
    resp, err := http.Post("http://localhost:8080/introduce", "text", buf)

    if err != nil {
    	t.Error("Was not able to comunicate with server")
    } else if resp.StatusCode != 400 {
    	t.Error("Server should return 400 error code")
    }
}

func TestMainPostBadArguments(t *testing.T) {
	buf := bytes.NewBufferString("{\"title\":\"Incorrect info!.\"}")
    resp, err := http.Post("http://localhost:8080/introduce", "text", buf)

    if err != nil {
    	t.Error("Was not able to comunicate with server")
    } else if resp.StatusCode != 400 {
    	t.Error("Server should return 400 error code")
    }
}

func TestMainPostMissingArguments(t *testing.T) {
	buf := bytes.NewBufferString("{\"firstname\":\"bob\"}")
    resp, err := http.Post("http://localhost:8080/introduce", "text", buf)

    if err != nil {
    	t.Error("Was not able to comunicate with server")
    } else if resp.StatusCode != 400 {
    	t.Error("Server should return 400 error code")
    }
}

func TestMainPostSuccess1(t *testing.T) {
	firstname, lastname := "bob", "rogers"
	buf := bytes.NewBufferString("{\"firstname\":\"" +firstname+ "\", \"lastname\":\"" +lastname+ "\"}")
    resp, err := http.Post("http://localhost:8080/introduce", "text", buf)

    if err != nil {
    	t.Error("Was not able to comunicate with server")
    } else if resp.StatusCode != 200 {
    	t.Error("Response should have code 200")
    } else {
    	bs, err := ioutil.ReadAll(resp.Body)
    	if err!= nil {
    		t.Error("Error reading response")
    	} else {
    		response := string(bs)
    		if !(strings.Contains(response, firstname) && strings.Contains(response , lastname)) {
    			t.Error("Response does not contain first and last name")
    		}
    	}
    }
}

func TestMainPostSuccess2(t *testing.T) {
	firstname, lastname := "jane", "jones"
	buf := bytes.NewBufferString("{\"firstname\":\"" +firstname+ "\", \"lastname\":\"" +lastname+ "\"}")
    resp, err := http.Post("http://localhost:8080/introduce", "text", buf)

    if err != nil {
    	t.Error("Was not able to comunicate with server")
    } else if resp.StatusCode != 200 {
    	t.Error("Response should have code 200")
    } else {
    	bs, err := ioutil.ReadAll(resp.Body)
    	if err!= nil {
    		t.Error("Error reading response")
    	} else {
    		response := string(bs)
    		if !(strings.Contains(response, firstname) && strings.Contains(response , lastname)) {
    			t.Error("Response does not contain first and last name")
    		}
    	}
    }
}