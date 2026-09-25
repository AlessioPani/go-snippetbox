FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/snippetbox ./cmd/web

FROM alpine:3.22 AS runtime

RUN addgroup -S -g 10001 app \
    && adduser -S -u 10001 -G app app \
    && mkdir -p /app/db-data /app/tls \
    && chown -R 10001:10001 /app

WORKDIR /app

COPY --from=builder --chown=10001:10001 /out/snippetbox ./snippetbox
COPY --from=builder --chown=10001:10001 /src/db-data/snippetbox.db ./db-data/snippetbox.db

USER 10001:10001

EXPOSE 8080

ENTRYPOINT ["./snippetbox"]

FROM alpine:3.22 AS certgen

RUN apk add --no-cache openssl \
    && addgroup -S -g 10001 app \
    && adduser -S -u 10001 -G app app \
    && mkdir -p /certs \
    && chown 10001:10001 /certs

USER 10001:10001

CMD ["/bin/sh", "-ec", "if [ ! -s /certs/cert.pem ] || [ ! -s /certs/key.pem ]; then openssl req -x509 -newkey rsa:2048 -nodes -keyout /certs/key.pem -out /certs/cert.pem -days 3650 -subj '/CN=localhost' -addext 'subjectAltName=DNS:localhost,IP:127.0.0.1'; chmod 600 /certs/key.pem; fi"]
