FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o sample-go-app main.go

FROM alpine:latest
COPY --from=builder /app/sample-go-app /usr/local/bin/sample-go-app
ENTRYPOINT ["/usr/local/bin/sample-go-app"]
