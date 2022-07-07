FROM golang:1.17-buster as builder

EXPOSE 8080

WORKDIR /app

COPY . /app
RUN go mod download

RUN go mod tidy

RUN go build -o /app/service ./server/main.go 

RUN chmod +x /app/service

FROM debian:buster-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

ENV GIN_MODE=release

COPY --from=builder /app/service /app/service

CMD ["/app/service"]