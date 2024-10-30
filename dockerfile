FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && CGO_ENABLED=0 go build -o alioth_app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/alioth_app .

RUN chmod +x alioth_app

EXPOSE 10000

ENTRYPOINT ["/app/alioth_app"]