FROM golang:latest AS builder
WORKDIR /src/app

COPY . .

RUN go build -o kline cmd/main.go 

FROM alpine:latest AS runner
COPY --from=builder /src/app/kline /usr/local/bin/

CMD [ "kline" , "version"]
