package server

import (
	"hackathon/db"
	"log"
	"net/http"
	"os"
)

// サーバーのセットアップと起動
func Start() error {
	// データベース接続の初期化
	dbConn, err := db.InitDB()
	if err != nil {
		return err
	}
	defer dbConn.Close()

	// ルーターの設定
	router := NewRouter(dbConn)

	// サーバーのポート設定
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("サーバーをポート%sで起動します...", port)
	return http.ListenAndServe(":"+port, router)
}
