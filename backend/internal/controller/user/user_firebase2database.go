package usercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (uc *UserController) GetUserIDByFirebaseUID(w http.ResponseWriter, r *http.Request) {
	// リクエストボディをパース
	firebaseUID := r.Header.Get("UserID")
	if firebaseUID == "" {
		http.Error(w, "UserID missing in request context", http.StatusUnauthorized)
		return
	}

	userID, err := uc.userUsecase.GetUserIDByFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user ID: %v", err), http.StatusInternalServerError)
		return
	}

	// レスポンスを返す
	resp := struct {
		UserID string `json:"user_id"`
	}{
		UserID: userID,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}