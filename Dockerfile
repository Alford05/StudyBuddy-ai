# 1. Build backend
FROM golang:1.22 AS builder
WORKDIR /app
COPY backend/ .
RUN go build -o server .

# 2. Serve frontend + backend
FROM alpine
WORKDIR /app

COPY --from=builder /app/server .
COPY frontend/ ./frontend/

EXPOSE 8080

CMD ["./server"]
