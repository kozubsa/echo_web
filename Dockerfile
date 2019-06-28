FROM golang:1.12.6-alpine as builder

ENV GO111MODULE on

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    go mod download

COPY . .

RUN go build -v

FROM alpine:3.9

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/web /app/app

CMD ["/app/app"]

EXPOSE 8080