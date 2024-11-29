package usercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (uc *UserController) UpdateUserPrivacy(w http.ResponseWriter, r *http.Request) {
	var request struct {
		IsPrivate bool `json:"is_private"`
	}
	ID, _, err := uc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), r)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエストの解析に失敗しました: %v", err), http.StatusBadRequest)
		return
	}

	user, err := uc.userUsecase.UpdateUserPrivacy(context.Background(), request.IsPrivate, ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザープロフィール更新に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}


}