FROM golang:1.23.8

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

EXPOSE 8090

CMD ["./main"]
