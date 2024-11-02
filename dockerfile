FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./pr-trigger-go ./cmd/main.go

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/pr-trigger-go .
EXPOSE 3000
ENTRYPOINT ["./pr-trigger-go"]