package authcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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