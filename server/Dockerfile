FROM golang as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux go build -o server

FROM ubuntu

WORKDIR /app
COPY --from=builder /app/server .

CMD ["./server"]