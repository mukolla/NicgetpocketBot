FROM golang:1.20.4-alpine3.18 as build

COPY . /github.com/mukolla/nicgetpocketbot/
WORKDIR /github.com/mukolla/nicgetpocketbot/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/mukolla/nicgetpocketbot/bin/bot .
COPY --from=0 /github.com/mukolla/nicgetpocketbot/configs configs/

EXPOSE 8183

CMD ["./bot"]