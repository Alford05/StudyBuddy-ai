# 1. Build the Go backend
FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY backend/ .
RUN go build -o server .

# 2. Build the React frontend
FROM node:18 AS react-builder
WORKDIR /frontend

COPY frontend/ .
RUN npm install
RUN npm run build

# 3. Serve frontend + backend with a single container
FROM golang:1.23
WORKDIR /app

# Install nginx
RUN apt-get update && apt-get install -y nginx && rm -rf /var/lib/apt/lists/*

# Copy Go server
COPY --from=builder /app/server .

# Copy React build
COPY --from=react-builder /frontend/dist /app/frontend/dist

# Copy Nginx config
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 8080
EXPOSE 80

CMD ["sh", "-c", "nginx && ./server"]


