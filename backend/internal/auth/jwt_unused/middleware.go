package jwtunused

import (
	"context"
	"net/http"
	"strings"
)

// ユーザーIDをコンテキストキーとして設定
type contextKey string

const UserIDKey contextKey = "userID"

// 認証ミドルウェア
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authorizationヘッダーからトークンを取得
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorizationヘッダーが必要です", http.StatusUnauthorized)
			return
		}

		// "Bearer {token}" の形式を分離
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			http.Error(w, "トークンの形式が正しくありません", http.StatusUnauthorized)
			return
		}

		// トークンを検証
		// claims, err := ValidateToken(token)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusUnauthorized)
		// 	return
		// }

		// // ユーザーIDをコンテキストに格納
		// ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		// next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// コンテキストからユーザーIDを取得
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDKey).(string)
	return userID, ok
}
