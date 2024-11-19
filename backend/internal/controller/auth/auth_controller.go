package authController

import (
	"encoding/json"
	"fmt"
	"hackathon/internal/usecase/user"
	"hackathon/internal/dao/user"
	"database/sql"
	"net/http"
	"hackathon/internal/auth"
	"time"
)

type AuthController struct {
	signInUsecase *usecase.UserSignInUsecase
}

func NewAuthController(dbConn *sql.DB) *AuthController {
	signInDAO := dao.NewUserSignInDAO(dbConn)
	signInUsecase := usecase.NewUserSignInUsecase(signInDAO)
	return &AuthController{signInUsecase: signInUsecase}
}

// サインインエンドポイント
func (ac *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// リクエストの解析
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエスト解析エラー: %v", err), http.StatusBadRequest)
		return
	}

	// 必須フィールドの検証
	if request.Username == "" || request.Password == "" {
		http.Error(w, "ユーザー名とパスワードが必要です", http.StatusBadRequest)
		return
	}

	// サインイン処理
	token, err := ac.signInUsecase.SignIn(r.Context(), request.Username, request.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("サインイン失敗: %v", err), http.StatusUnauthorized)
		return
	}
	// 成功したらログに出力
	fmt.Printf("サインイン成功: %s\n", request.Username)

	// トークンを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

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