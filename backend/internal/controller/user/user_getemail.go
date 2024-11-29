package usercontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


// GetUserEmailByUsernameHandler はユーザー名からメールアドレスを取得するAPIハンドラー
func (uc *UserController) GetUserEmailByUsername(w http.ResponseWriter, r *http.Request) {
	// URLパスからパラメータ「username」を取得
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		http.Error(w, "usernameパラメータが指定されていません", http.StatusBadRequest)
		return
	}

	// usernameをログに出力
	fmt.Println("username: ", username)

	// ユースケース層のロジックを呼び出す
	ctx := r.Context()
	email, err := uc.userUsecase.GetEmailByUsernameUsecase(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 結果をJSONとしてレスポンス
	response := map[string]string{"email": email}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
