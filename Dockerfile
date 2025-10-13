FROM golang:1.25.2-alpine3.22 AS builder

WORKDIR /app

COPY ./src/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o shutdown_server shutdown_server.go


FROM selenium/standalone-chrome:141.0

USER root

COPY --from=builder /app/shutdown_server /opt/bin/shutdown_server

COPY ./src/entry_point.sh /opt/bin/

RUN chmod +x /opt/bin/entry_point.sh

USER seluser

CMD ["/opt/bin/entry_point.sh"]
