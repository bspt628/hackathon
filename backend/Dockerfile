# ビルドステージ
FROM golang:1.21 AS build

WORKDIR /app

# 現在のディレクトリ（backend）の内容をコピー
COPY . .

# 依存関係をダウンロード
RUN go mod download

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# 実行ステージ
FROM debian:stable-slim AS app

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# ビルドしたバイナリをコピー
COPY --from=build /app/app .

# 設定ファイルと証明書をコピー（必要な場合）
COPY config ./config
COPY cert ./cert

# ポート8080を公開（必要に応じて変更）
EXPOSE 8080

# アプリケーションを実行
CMD ["./app"]