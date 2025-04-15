# ============ BASE ===========
FROM golang:1.24.1-alpine3.20 AS base

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# ========= BUILDER ==========
FROM base AS builder

WORKDIR /app

COPY . .

# RUN go mod download

RUN go build -o backend .

# ========= RUNNER ==========
FROM golang:1.24.1-alpine3.20 AS release

WORKDIR /home/node

COPY --from=builder /app/backend ./backend

EXPOSE $APP_PORT

CMD [ "./backend" ]