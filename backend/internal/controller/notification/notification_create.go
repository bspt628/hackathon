package notificationcontroller

import (
	"net/http"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hackathon/internal/model"
)

func (nc *NotificationController) CreateNotifications(w http.ResponseWriter, r *http.Request) {
	// FirebaseAuthMiddleware で設定された UserID を取得
	firebaseUID := r.Header.Get("UserID")
	if firebaseUID == "" {
		http.Error(w, "UserID missing in request context", http.StatusUnauthorized)
		return
	}

	userID, err := nc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("いいねするユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	var request struct {
		Type string `json:"type"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエストの解析に失敗しました: %v", err), http.StatusBadRequest)
		return
	}

	if request.Type == "" {
		http.Error(w, "Type is required", http.StatusBadRequest)
		return
	}

	if request.Message == "" {
		http.Error(w, "Message is required", http.StatusBadRequest)
		return
	}

	params := model.CreateNotificationParams{
		ID: uuid.New().String(),
		UserID: userID,
		Type: request.Type,
		Message: request.Message,
	}

	err = nc.notificationUsecase.CreateNotifications(context.Background(), params)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create notification: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Notification created successfully"}`))
}


