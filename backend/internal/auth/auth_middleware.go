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

        // Firebaseの認証を使ってユーザー情報を取得
        user, err := GetUserInfo(token)
        if err != nil {
            http.Error(w, fmt.Sprintf("Firebase authentication failed: %v", err), http.StatusUnauthorized)
            return
        }

        // ユーザー情報をリクエストにセットするなどの処理
        r.Header.Set("UserID", user.UID)

        next.ServeHTTP(w, r)
    })
}
