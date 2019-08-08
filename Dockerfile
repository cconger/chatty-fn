FROM golang:1-alpine as builder

WORKDIR /usr/app
COPY . .
RUN go build -mod=vendor -o function .

FROM openfaas/of-watchdog:0.5.3 as watchdog

FROM alpine:latest

RUN apk --no-cache add ca-certificates \
    && addgroup -S app && adduser -S -g app app \
    && mkdir -p /home/app \
    && chown app /home/app

WORKDIR /home/app

COPY --from=builder /usr/app/function .
COPY --from=watchdog /fwatchdog .

RUN chown -R app /home/app

USER app

ENV fprocess="/home/app/function"
ENV mode="http"
ENV upstream_url="http://127.0.0.1:5000"

EXPOSE 8080

CMD ["./fwatchdog"]
