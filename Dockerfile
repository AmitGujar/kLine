FROM golang:latest AS builder
WORKDIR /src/app

COPY . .

RUN go build -o kline cmd/main.go 

FROM ubuntu:noble AS runner

COPY --from=builder /src/app/kline /usr/local/bin/
RUN chmod 777 /usr/local/bin/kline

ENTRYPOINT [ "kline" ]

CMD ["--version"]
