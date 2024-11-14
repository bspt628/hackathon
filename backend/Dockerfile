FROM golang:1.22.3

# 作業ディレクトリを指定
WORKDIR /go/src/backend

# 必要なパッケージをインストールする
RUN apt-get update && apt-get install -y git

# ホストのGoコードをコンテナ内にコピー
COPY . .

# 必要なGoの依存関係をインストール
RUN go mod tidy

# アプリケーションを実行
CMD ["go", "run", "main.go"]
