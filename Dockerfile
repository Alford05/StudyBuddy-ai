# 1. Build the Go backend
FROM golang:1.22 AS builder
WORKDIR /app

# Copy Go files and build the Go server
COPY backend/ .
RUN go build -o server .

# 2. Build the React frontend
FROM node:18 AS react-builder
WORKDIR /frontend

# Copy React app files and install dependencies
COPY frontend/ .
RUN npm install
RUN npm run build

# 3. Serve frontend + backend with a single container
FROM alpine:latest
WORKDIR /app

# Copy the Go server from the builder stage
COPY --from=builder /app/server .

# Copy the React build from the react-builder stage
COPY --from=react-builder /frontend/build /app/frontend/build

# Install necessary packages (nginx, or any web server for serving frontend)
RUN apk --no-cache add nginx

# Expose ports
EXPOSE 8080
EXPOSE 80

# Configure Nginx to serve the React frontend
COPY nginx.conf /etc/nginx/nginx.conf

# Command to run both Go server and Nginx
CMD ["sh", "-c", "nginx && ./server"]
