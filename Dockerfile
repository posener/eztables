FROM golang:1.9.2-alpine as builder
WORKDIR /go/src/github.com/posener/eztables
COPY . .
RUN apk add --no-cache git
RUN go get -u ./...
RUN go build

FROM alpine:3.7
RUN apk add --no-cache iptables
COPY --from=builder /go/src/github.com/posener/eztables/eztables /usr/bin/eztables
ENTRYPOINT ["eztables"]
