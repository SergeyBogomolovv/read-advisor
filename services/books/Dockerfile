FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./books cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache libc6-compat

COPY --from=builder /app/books ./books

CMD [ "./books" ]