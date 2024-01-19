FROM golang:1.21.3-alpine3.17 AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go build -o api ./cmd/api

FROM alpine:3.17 AS binary
COPY --from=base /app/api .
EXPOSE 8080
CMD ["./api"]