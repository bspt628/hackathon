package usercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"hackathon/domain"
)

func (uc *UserController) UpdateUserNotifications(w http.ResponseWriter, r *http.Request) {
	var request struct {
		NotificationSettings domain.NotificationSettings `json:"notification_settings"`
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

	// 入力バリデーション
	if request.NotificationSettings.Frequency == "" {
		http.Error(w, "通知頻度を指定してください", http.StatusBadRequest)
		return
	}

	// ユースケースの呼び出し
	user, err := uc.userUsecase.UpdateUserNotifications(context.Background(), request.NotificationSettings, ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザープロフィール更新に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// レスポンスのエンコード
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}
