from golang:1.19 as golang

workdir /app

COPY go.mod .
COPY server.go .

RUN go build -o server

ENTRYPOINT ["./server"]

CMD []
