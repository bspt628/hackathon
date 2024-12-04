package auth

import (
	"net/http"
	"fmt"
	"log"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// corsミドルウェアを使用することをコンソールに出力
		fmt.Println("CORS middleware in use")
		// 必要な CORS ヘッダーを設定
		w.Header().Set("Access-Control-Allow-Origin", "https://hackathon-five-rho.vercel.app") // "*" は任意のオリジンを許可
		w.Header().Set("Access-Control-Allow-Credentials", "true") // クッキーを許可
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // 許可する HTTP メソッド
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // 許可するヘッダー

		// プリフライトリクエスト (OPTIONS) に対して早期に応答
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}
		// ヘッダーをコンソールに出力
		fmt.Println("Request Headers: ", r.Header)
		log.Println("Request Headers: ", r.Header)

		next.ServeHTTP(w, r)
	})
}