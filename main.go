package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/go-krd/meetup-1-workshop/pkg/brute"
)

func main() {
	const hash = "95ebc3c7b3b9f1d2c40fec14415d3cb8"

	before := time.Now()

	pass, err := brute.Multi(hash, runtime.NumCPU())
	if err != nil {
		log.Fatalf("failed to bruteforce: %s", err)
	}

	took := time.Now().Sub(before)

	fmt.Printf("Password found: %s\n", pass)
	fmt.Printf("Took: %s", took.String())
}
