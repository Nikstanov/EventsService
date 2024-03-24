FROM golang:alpine AS builder

WORKDIR /build

COPY .. .

RUN go build -o eventsService ./cmd/eventsservice/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /build/eventsService /app/eventsService

CMD ["./eventsService"]