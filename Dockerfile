
FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /flashbook .  # Simpan binary di root


FROM alpine:latest

RUN apk add --no-cache ca-certificates libc6-compat

WORKDIR /app


COPY --from=builder /flashbook /flashbook


COPY .env ./

EXPOSE 8080


CMD ["/flashbook"]