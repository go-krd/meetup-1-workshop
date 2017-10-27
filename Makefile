pkg := github.com/go-krd/meetup-1-workshop
port := 3000

imageName = go-krd/bruteforcer
imageTag = latest

build:
	@go build -o bin/bruteforcer $(pkg)/cmd/bruteforcer

run:
	@SERVICE_PORT=$(port) bin/bruteforcer

test:
	@go test -v -cover ./...

image:
	@docker image build -t $(imageName):$(imageTag) .