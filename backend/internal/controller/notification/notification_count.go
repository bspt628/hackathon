package notificationcontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (nc *NotificationController) CountUnreadNotifications(w http.ResponseWriter, r *http.Request) {
	// ユーザIDを取得
	firebaseUID := r.Header.Get("UserID")

	userID, err := nc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	count, err := nc.notificationUsecase.CountUnreadNotifications(context.Background(), userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("未読の通知のカウントに失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]int64{"your Unread Notifications": count}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
	
}