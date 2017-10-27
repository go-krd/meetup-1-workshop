FROM golang:1.9-alpine as build

WORKDIR /go/src/github.com/go-krd/meetup-1-workshop

COPY . .

RUN go test -v -cover ./...
RUN go build -o bin/bruteforcer ./cmd/bruteforcer

FROM alpine:3.6

COPY --from=build /go/src/github.com/go-krd/meetup-1-workshop/bin/bruteforcer /service/bruteforcer

ENV SERVICE_PORT 3000

ENTRYPOINT ["/service/bruteforcer"]