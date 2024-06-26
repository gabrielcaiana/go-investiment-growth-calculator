# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build the binary.
COPY . .
RUN CGO_ENABLED=0 go build -o main .

# Final stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
