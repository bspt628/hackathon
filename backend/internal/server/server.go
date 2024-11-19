package server

import (
	"hackathon/db"
	"log"
	"net/http"
	"os"
	"github.com/rs/cors"
	"strings"
	"hackathon/internal/auth"
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

	log.Printf("サーバーをポート%sで起動します...", port)

	// verifyTokenHandler を追加
	http.HandleFunc("/verify-token", verifyTokenHandler)

	return http.ListenAndServe(":"+port, handler)
}


// verifyTokenHandlerは、Firebase認証のIDトークンを検証するためのハンドラーです。
func verifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	// AuthorizationヘッダーからIDトークンを取得
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// "Bearer "を取り除いてトークン部分を取得
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// FirebaseのIDトークンを検証し、UIDを取得
	uid, err := auth.VerifyIDToken(token)
	if err != nil {
		http.Error(w, "Firebase authentication failed: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// UIDが正常に取得できた場合、成功メッセージを返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Token is valid", "uid": "` + uid + `"}`))
}