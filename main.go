package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/go-krd/meetup-1-workshop/pkg/brute"
)

func main() {
	const hash = "95ebc3c7b3b9f1d2c40fec14415d3cb8"

	before := time.Now()

	pass, err := bruteforce(hash)
	if err != nil {
		log.Fatalf("failed to bruteforce: %s", err)
	}

	took := time.Now().Sub(before)

	fmt.Printf("Password found: %s\n", pass)
	fmt.Printf("Took: %s", took.String())
}

func bruteforce(hash string) (password string, err error) {
	b, err := hex.DecodeString(hash)
	if err != nil {
		return "", fmt.Errorf("decode hex: %s", err)
	}

	var target [md5.Size]byte
	copy(target[:], b)

	source := []byte("00000")
	for md5.Sum(source) != target {
		brute.NextWord(source)

		if bytes.Equal(source, []byte("00000")) {
			return "", fmt.Errorf("password not found")
		}
	}

	return string(source), nil
}
