# Builder
FROM golang:1.17.1-alpine3.14 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

COPY . /app
WORKDIR /app   

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

COPY --from=builder /app/configs configs
COPY --from=builder /app/.bin/app app

CMD ["./app"]