FROM golang:1.23.4-alpine AS builder

WORKDIR /onlinesong

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY .env .

RUN go build -o onlinesong .

FROM alpine:latest

WORKDIR /onlinesong

COPY --from=builder /onlinesong/onlinesong .
COPY --from=builder /onlinesong/.env .
COPY --from=builder /onlinesong/logs ./logs
COPY --from=builder /onlinesong/docs .

EXPOSE 8080

CMD ["./onlinesong"]