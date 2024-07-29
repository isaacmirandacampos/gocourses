FROM golang:1.22 AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin/app

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /go/bin/app /app/

EXPOSE 8080

CMD ["/app/app"]
