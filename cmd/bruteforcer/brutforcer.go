package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/go-krd/meetup-1-workshop/pkg/brute"
)

const (
	envServicePort = "SERVICE_PORT"
)

func main() {
	port := os.Getenv(envServicePort)
	if port == "" {
		panic("environment variable SERVICE_PORT is empty")
	}

	addr := fmt.Sprintf(":%s", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/brute", bruteHandler)

	log.Printf("Listen %s", addr)
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}

// try: curl /brute?hash=95ebc3c7b3b9f1d2c40fec14415d3cb8
func bruteHandler(w http.ResponseWriter, req *http.Request) {
	hash := req.URL.Query().Get("hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "hash is not specified")
		return
	}

	before := time.Now()
	pass, err := brute.Multi(hash, runtime.NumCPU())
	took := time.Now().Sub(before)

	w.Header().Set("X-Duration", took.String())

	if err != nil {
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, pass)
	}
}
