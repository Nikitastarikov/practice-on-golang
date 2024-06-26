FROM golang:1.22 AS builder

WORKDIR /src
COPY . .

RUN go build -a -installsuffix cgo -o  server/server -ldflags "-s -w" ./server/server.go

CMD ["./server/server"]

EXPOSE 5300
