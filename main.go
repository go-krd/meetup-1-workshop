package main

import "fmt"

func main() {
	const hash = "95ebc3c7b3b9f1d2c40fec14415d3cb8"
	pass := bruteforce(hash)
	fmt.Println(pass)
}

func bruteforce(hash string) (password string) {
	panic("implement me")
}