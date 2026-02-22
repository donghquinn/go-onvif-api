# ============ BASE ===========
FROM golang:1.25.5-alpine3.22 AS base

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# ========= BUILDER ==========
FROM base AS builder

# Install security updates
RUN apk update && apk upgrade && apk add --no-cache ca-certificates

WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build with security flags
RUN go build -ldflags="-w -s -extldflags '-static'" -a -installsuffix cgo -o backend cmd/main.go

# ========= RUNNER ==========
FROM alpine:3.22 AS release

# Build arguments for flexible UID/GID
ARG USER_ID=1000
ARG GROUP_ID=1000

# Install security updates and essential packages
RUN apk update && apk upgrade && \
    apk add --no-cache ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

# Create uploader user and group with configurable UID/GID
RUN addgroup -g ${GROUP_ID} -S node && \
    adduser -u ${USER_ID} -S node -G node -h /home/node -s /sbin/nologin

# Set working directory
WORKDIR /home/node

# Copy binary and set ownership
COPY --from=builder --chown=node:node /app/backend ./backend

# Create necessary directories with proper permissions BEFORE switching user
RUN mkdir -p logs tmp node && \
    chown -R node:node /home/node logs tmp && \
    chmod 755 /home/node && \
    chmod 775 logs tmp && \
    chmod 755 node

# Switch to non-root user
USER node

# Use exec form for better signal handling
CMD ["./backend"]