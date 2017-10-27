pkg := github.com/go-krd/meetup-1-workshop
port := 3000

build:
	@go build -o bin/bruteforcer $(pkg)/cmd/bruteforcer

run:
	@SERVICE_PORT=$(port) bin/bruteforcer