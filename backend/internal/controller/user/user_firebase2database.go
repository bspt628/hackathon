package usercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (uc *UserController) GetUserIDByFirebaseUID(w http.ResponseWriter, r *http.Request) {
	// リクエストボディをパース
	type Request struct {
		FirebaseUID string `json:"firebase_uid"`
	}
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request body: %v", err), http.StatusBadRequest)
		return
	}

	// ユーザーIDを取得
	userID, err := uc.userUsecase.GetUserIDByFirebaseUID(context.Background(), req.FirebaseUID)
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