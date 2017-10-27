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
	io.WriteString(w, "Hello, bruteforcer!")
}
