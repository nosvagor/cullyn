# build stage
FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -o main cmd/main.go

# run stage
FROM alpine:3.19
RUN apk add --no-cache tzdata curl wget
WORKDIR /app
COPY --from=builder /app/main .
COPY static/ /app/static/
ENV GIN_MODE=release
EXPOSE 3000
CMD ["./main"]
