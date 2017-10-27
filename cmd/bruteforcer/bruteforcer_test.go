package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBruteHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/brute", bruteHandler)

	srv := httptest.NewServer(mux)

	client := srv.Client()

	resp, err := client.Get(srv.URL + "/brute?hash=95ebc3c7b3b9f1d2c40fec14415d3cb8")
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}

	expected := "zzzzz"
	if string(b) != expected {
		t.Fatalf("Expected result to be %s but got %s", expected, string(b))
	}

}
