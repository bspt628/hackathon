package main

import (
	// "github.com/joho/godotenv"
	"hackathon/internal/auth"
	"hackathon/internal/server"
	"log"
)

func main() {
	// Firebaseの初期化
    err := auth.InitFirebase()
    if err != nil {
        log.Fatalf("Firebase initialization failed: %v", err)
    }
	
	// サーバーのセットアップと起動
	if err := server.Start(); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
