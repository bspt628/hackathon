package authcontroller

import (
	"fmt"
	"net/http"
	"database/sql"
	"hackathon/internal/dao/user"
	"hackathon/internal/usecase/user"
	"encoding/json"
	"hackathon/db/sqlc/generated"
)

type PasswordResetController struct {
	passwordResetUsecase *userusecase.UserPasswordResetUsecase
}

func NewPasswordResetController(dbConn *sql.DB) *PasswordResetController {
	queries := sqlc.New(dbConn)
	passwordResetDAO := userdao.NewUserPasswordResetDAO(queries)
	passwordResetUsecase := userusecase.NewUserPasswordResetUsecase(passwordResetDAO)
	return &PasswordResetController{passwordResetUsecase: passwordResetUsecase}
}

// パスワードリセットリクエスト (トークン送信)
func (prc *PasswordResetController) HandlePasswordResetRequest(w http.ResponseWriter, r *http.Request) {
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

	// パスワードリセット処理
	err := prc.passwordResetUsecase.RequestPasswordReset(r.Context(), req.Email)
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
func (prc *PasswordResetController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}

	// リクエスト解析
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("リクエスト解析エラー: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println(req)

	// トークンとパスワードのバリデーション
	if req.Token == "" || req.Password == "" {
		http.Error(w, "トークンとパスワードは必須です", http.StatusBadRequest)
		return
	}

	// パスワード更新処理
	err := prc.passwordResetUsecase.ResetPassword(r.Context(), req.Token, req.Password)
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
