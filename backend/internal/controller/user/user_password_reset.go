package usercontroller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"context"
)

// パスワードリセットリクエスト (トークン送信)
func (uc *UserController) PasswordResetRequest(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	// リクエスト解析
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "無効なリクエスト", http.StatusBadRequest)
		return
	}

	// メールアドレスが必要かどうかのバリデーション
	if req.Email == "" {
		http.Error(w, "メールアドレスは必須です", http.StatusBadRequest)
		return
	}

	ID, _, err := uc.GetUserIDFromFirebaseUID(context.Background(), r)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// ユーザーIDからメールアドレスを取得
	user, err := uc.userUsecase.GetUser(context.Background(), ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザー情報の取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// 取得したメールアドレスが空の場合はエラー
	if user.Email == "" {
		http.Error(w, "メールアドレスが見つかりません", http.StatusInternalServerError)
		return
	}
	// 取得したメールアドレスがリクエストのメールアドレスと異なる場合はエラー
	if req.Email != user.Email {
		http.Error(w, "メールアドレスが一致しません", http.StatusBadRequest)
		return
	}

	// パスワードリセット処理
	err = uc.userUsecase.RequestPasswordReset(r.Context(), req.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("パスワードリセットに失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "パスワードリセットメールを送信しました",
	})
}

// パスワード変更 (トークン検証とパスワード更新)
func (uc *UserController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Password string `json:"password"`
	}

	// リクエスト解析
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("リクエスト解析エラー: %v", err), http.StatusBadRequest)
		return
	}

	// パスパラメータからトークンを取得
	token := r.URL.Query().Get("token")

	// トークンとパスワードのバリデーション
	if token == "" || req.Password == "" {
		http.Error(w, "トークンとパスワードは必須です", http.StatusBadRequest)
		return
	}

	// パスワード更新処理
	err := uc.userUsecase.ResetPassword(r.Context(), token, req.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("パスワード更新失敗: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "パスワードが更新されました",
	})
}
