package usercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"hackathon/internal/model"
)

// CreateUser は新規ユーザーを作成するエンドポイント
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://hackathon-five-rho.vercel.app") // "*" は任意のオリジンを許可
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // 許可する HTTP メソッド
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // 許可するヘッダー
	// リクエストボディからユーザー情報を取得
	var request model.UserCreateRequest

	// リクエストのJSONデータを構造体にバインド
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエストの解析に失敗しました: %v", err), http.StatusBadRequest)
		return
	}

	// 必須フィールドのバリデーション
	if request.Email == "" || request.Password == "" || request.Username == "" {
		http.Error(w, "必須フィールドが不足しています", http.StatusBadRequest)
		return
	}

	// 新規ユーザーを作成
	user, err := uc.userUsecase.CreateUser(context.Background(),request)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザー作成に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// JSON形式でレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}