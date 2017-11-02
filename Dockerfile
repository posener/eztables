FROM golang:1.9.2-alpine
RUN apk add --no-cache git iptables
WORKDIR /go/src/github.com/posener/eztables
COPY . .
RUN go install
CMD ["eztables"]
