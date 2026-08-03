FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o movies-service .

FROM alpine:latest

RUN apk --no-cache add ca-certificates wget

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/movies-service .

# Expose the application port
EXPOSE 8081

# Command to run the executable
CMD ["./movies-service"]