# meetup-1-workshop

A simple demonstrational bruteforce service.

## Run

#### Local

    SERVICE_PORT=3000 make build run

#### Docker

    make image
    docker run --rm -p 3000:3000 go-krd/bruteforcer:latest