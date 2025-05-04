FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go

CMD ["sh", "-c", "sleep 10 && ./main"]
