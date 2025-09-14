# Use official Go image for building
FROM golang:1.22 as builder

WORKDIR /app
COPY . .

# Build the Go binary
RUN go mod tidy
RUN go build -o server main.go

# Use a lightweight image for running
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/server .
COPY static ./static

# Expose Railway's PORT
EXPOSE 8080

CMD ["./server"]