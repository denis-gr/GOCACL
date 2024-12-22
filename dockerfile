FROM golang:1.23.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN mkdir -p build && go build -o build/ ./...

FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache libc6-compat
EXPOSE 8080
COPY --from=builder /app/build/* .
CMD ["./server"]
