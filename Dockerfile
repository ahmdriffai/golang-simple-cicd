# Stage 1: Build
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum (jika ada)
COPY go.mod ./

# Download dependencies (cache layer ini untuk efisiensi build)
RUN go mod download

COPY . .

# Build aplikasi
RUN go build -o main .

# Stage 2: Minimal runtime
FROM alpine:latest

WORKDIR /app

# Copy executable dari stage build
COPY --from=builder /app/main .

# Expose port yang digunakan oleh aplikasi
EXPOSE 7000

# Jalankan aplikasi
CMD ["./main"]