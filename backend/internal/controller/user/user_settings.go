package usercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (uc *UserController) UpdateUserSettings(w http.ResponseWriter, r *http.Request) {
	var request struct {
		DisplayName string `json:"display_name"`
		BirthDate   string `json:"birth_date"`
		Language    string `json:"language"`
	}
	ID, _, err := uc.GetUserIDFromFirebaseUID(context.Background(), r)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエストの解析に失敗しました: %v", err), http.StatusBadRequest)
		return
	}

	if request.DisplayName == "" && request.BirthDate == "" && request.Language == "" {
		http.Error(w, "何か変更をしてください", http.StatusBadRequest)
		return
	}

	user, err := uc.userUsecase.UpdateUserSettings(context.Background(), request.DisplayName, request.BirthDate, request.Language, ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザー設定更新に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}

}
