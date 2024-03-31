# Stage 1: Build stage
FROM golang:1.22.1-alpine as build-base

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and download all dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the entire source code
COPY . .

# Test the application
RUN CGO_ENABLED=0 go test -v

# Build the application
RUN go build -o ./out/fun-exercise-api .

# ===================================================
# Stage 2: Final stage
FROM alpine:3.16.2

# Copy the built binary and .env from the previous stage
COPY --from=build-base /app/out/fun-exercise-api /app/fun-exercise-api
# COPY --from=build-base /app/out/.env /app/.env

# Set the binary as the entry point
CMD ["/app/fun-exercise-api"]
