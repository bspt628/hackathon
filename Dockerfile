# ビルドステージ
FROM golang:1.22.3 AS build
WORKDIR /app
# バックエンドのソースコードをコピー
COPY backend/ ./backend/
# 作業ディレクトリをbackendに変更
WORKDIR /app/backend
# 依存関係をダウンロード
RUN go mod download
# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o app .
# 実行ステージ
FROM debian:stable-slim AS app
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
# ビルドしたバイナリをコピー
COPY --from=build /app/backend/app .
# 設定ファイルと証明書をコピー
COPY backend/config ./config
COPY backend/cert ./cert
# ポート8080を公開（必要に応じて変更）
EXPOSE 8080
# アプリケーションを実行
CMD ["./app"]