package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	err := http.ListenAndServe(addr, http.HandlerFunc(helloHandler))
	log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	hash := req.URL.Query().Get("hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "hash is not specified")
		return
	}

	io.WriteString(w, "Hello, bruteforcer!")
}
