package utils

import (
	"context"
	"fmt"
	"net/http"
	"hackathon/internal/usecase/user"
)

// GetUserIDFromFirebaseUID は、リクエストヘッダから Firebase UID を取得し、
// Firebase UID に基づいてユーザーIDを取得します。
// 失敗した場合はエラーメッセージを返します。
func GetUserIDFromFirebaseUID(r *http.Request, userUsecase userusecase.UserUsecase) (string, error) {
	// リクエストヘッダから UserID を取得
	firebaseUID := r.Header.Get("UserID")
	if firebaseUID == "" {
		return "", fmt.Errorf("UserID missing in request context")
	}

	// Firebase UID に基づいてユーザーIDを取得
	id, err := userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		return "", fmt.Errorf("ユーザーIDの取得に失敗しました: %v", err)
	}

	return id, nil
}
