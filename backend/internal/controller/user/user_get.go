package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetUserByID はユーザーIDを指定してユーザー情報を取得するエンドポイント
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// ぱすぱらめーたからユーザーIDを取得
	vars := mux.Vars(r)
	userID := vars["id"]
	// クエリパラメータからユーザーIDを取得
	// userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "IDパラメータが指定されていません", http.StatusBadRequest)
		return
	}

	// コンテキストとともにユーザーを取得
	user, err := uc.userUsecase.GetUser(context.Background(), userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// JSON形式でレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}