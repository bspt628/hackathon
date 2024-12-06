package notificationcontroller

import (
	"net/http"
	"fmt"
	"context"
	"github.com/gorilla/mux"
)

func (nc *NotificationController) MarkNotificationsAsRead(w http.ResponseWriter, r *http.Request) {
	// パスパラメータから通知自体のIDを取得
	vars := mux.Vars(r)
	notificationID := vars["id"]


	err := nc.notificationUsecase.MarkNotificationsAsRead(context.Background(), notificationID)
	if err != nil {
		http.Error(w, fmt.Sprintf("通知の既読化に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Notifications marked as read"}`))
}