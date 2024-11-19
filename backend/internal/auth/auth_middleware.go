// ミドルウェアとして、リクエストに含まれる Authorization ヘッダーから ID トークンを取り出し、それを Firebase で検証
package auth

import (
	"fmt"
	"net/http"
	"strings"
)

// Firebase認証を行うミドルウェア
func FirebaseAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// リクエストのヘッダーからAuthorizationトークンを取得
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// "Bearer "の部分を削除してトークンを取得
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// FirebaseのIDトークンを検証してUIDを取得
		uid, err := VerifyIDToken(token)
		if err != nil {
			http.Error(w, fmt.Sprintf("Firebase authentication failed: %v", err), http.StatusUnauthorized)
			return
		}

		// ユーザーIDをリクエストにセット
		r.Header.Set("UserID", uid)

		// 次のハンドラーを呼び出す
		next.ServeHTTP(w, r)
	})
}
