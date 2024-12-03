package main

import (
	"github.com/joho/godotenv"
	"hackathon/internal/server"
	"log"
	"hackathon/internal/auth"
)

func main() {
	// 環境変数の読み込み
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf(".env ファイルの読み込みに失敗しました: %v", err)
	// }

	// Firebaseの初期化
    err = auth.InitFirebase()
    if err != nil {
        log.Fatalf("Firebase initialization failed: %v", err)
    }

	// サーバーのセットアップと起動
	if err := server.Start(); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
