# Stage 1: Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o lorcana ./cmd

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/lorcana .
EXPOSE 8000
CMD ["./lorcana"]
