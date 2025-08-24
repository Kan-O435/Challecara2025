# Goアプリケーションのビルドステージ
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Goモジュールをキャッシュするために、go.modとgo.sumをコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコンテナにコピー
COPY . .

# アプリケーションをビルド
RUN go build -o challecara2025-back .

# 本番環境用の最終イメージ
FROM alpine:latest

WORKDIR /root/

# ビルドステージからビルド済みのバイナリをコピー
COPY --from=builder /app/challecara2025-back .

# 実行ファイルを実行
CMD ["./challecara2025-back"]
