# Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o /app/bin/hello_world ./cmd/hello_world

EXPOSE 8080

ENTRYPOINT ["/app/bin/hello_world"]
