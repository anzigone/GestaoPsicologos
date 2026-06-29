FROM golang:1.22-alpine AS builder
WORKDIR /app
RUN apk --no-cache add git
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6
COPY . .
RUN swag init -g cmd/server/main.go --output docs
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/server /app/server
RUN mkdir -p /app/data && chmod 777 /app/data
EXPOSE 8080
VOLUME ["/app/data"]
CMD ["/app/server"]
