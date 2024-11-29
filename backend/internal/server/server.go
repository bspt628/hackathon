package server

import (
	"hackathon/db"
	"log"
	"net/http"
	"os"
	"github.com/rs/cors"
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

	// CORSミドルウェアを設定
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://your-frontend-domain.com", "http://127.0.0.1:3000"}, // 許可するオリジン
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// CORSミドルウェアを適用
	handler := c.Handler(router)

	// サーバーのポート設定
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("サーバーをポート%sで起動...", port)

	return http.ListenAndServe(":"+port, handler)
}


