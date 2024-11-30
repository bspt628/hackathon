package usercontroller

import (
	"context"
	"fmt"
	"net/http"
)

// GetUserIDFromFirebaseUID は、リクエストヘッダから Firebase UID を取得し、
// Firebase UID に基づいてユーザーIDを取得します。
// 失敗した場合はエラーメッセージを返します。
func (uc *UserController) GetUserIDFromFirebaseUID(ctx context.Context, r *http.Request) (string, string, error) {
	// リクエストヘッダから UserID を取得
	firebaseUID := r.Header.Get("UserID")
	if firebaseUID == "" {
		return "", "", fmt.Errorf("UserID missing in request context")
	}

	// Firebase UID に基づいてユーザーIDを取得
	id, err := uc.userUsecase.GetUserIDFromFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return "", "", fmt.Errorf("ユーザーIDの取得に失敗しました: %v", err)
	}
	fmt.Println("id[", id, "でログイン中")

	return id, firebaseUID, nil
}
