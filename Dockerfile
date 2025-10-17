FROM golang:1.23-alpine AS builder

WORKDIR /app

# Go toolchainの自動ダウンロードを有効化
ENV GOTOOLCHAIN=auto

# Goモジュールをキャッシュするために、go.modとgo.sumをコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコンテナにコピー
COPY . .

# アプリケーションをビルド
RUN go build -o challecara2025-back ./cmd/api

# 本番環境用の最終イメージ
FROM alpine:latest

WORKDIR /root/

# ビルドしたバイナリをコピー
COPY --from=builder /app/challecara2025-back .

EXPOSE 8080

CMD ["./challecara2025-back"]