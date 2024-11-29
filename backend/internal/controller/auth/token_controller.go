package authcontroller

import (
	"encoding/json"
	"fmt"
	"hackathon/internal/auth"
	"net/http"
	"time"
)


// トークン生成エンドポイント
func (ac *AuthController) GenerateToken(w http.ResponseWriter, r *http.Request) {
	// ユーザー認証後に発行（例としてリクエストからuser_idを取得）
	var request struct {
		UserID string `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエスト解析エラー: %v", err), http.StatusBadRequest)
		return
	}

	if request.UserID == "" {
		http.Error(w, "ユーザーIDが必要です", http.StatusBadRequest)
		return
	}

	// トークンを生成
	token, err := auth.GenerateToken(request.UserID, time.Hour*24)
	if err != nil {
		http.Error(w, fmt.Sprintf("トークン生成エラー: %v", err), http.StatusInternalServerError)
		return
	}

	// トークンを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}