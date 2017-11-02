FROM golang:1.9.2-alpine
RUN apk add --no-cache git iptables
WORKDIR /go/src/eztables
COPY . .
RUN go-wrapper download
RUN go-wrapper install
CMD ["eztables"]
