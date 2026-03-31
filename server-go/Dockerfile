FROM golang:1.25-alpine AS builder

WORKDIR /src

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /gopherAI .

FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# loadenv 必须能打开 .env；空文件表示变量仅来自容器环境（env_file / environment）
RUN touch /app/.env

COPY --from=builder /gopherAI /app/gopherAI

EXPOSE 3000

ENV APP_HOST=0.0.0.0 \
    APP_PORT=3000

CMD ["/app/gopherAI"]
