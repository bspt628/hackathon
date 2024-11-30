package followcontroller

import (
	"context"
	"fmt"
	"net/http"
	"encoding/json"
)


func (fc *FollowController) GetFollowers(w http.ResponseWriter, r *http.Request) {
	// FirebaseAuthMiddleware で設定された UserID を取得
	firebaseUID := r.Header.Get("UserID")
	if firebaseUID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return 
	}

	ID, err := fc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// Usecase層でフォロワー情報を取得
	followers, err := fc.followUsecase.GetFollowers(context.Background(), ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get followers: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"followers": followers,
	}
	
	// フォロワーがいない場合にメッセージを追加
	if len(followers) == 0 {
		response["followers"] = []interface{}{}
		response["message"] = "This user has no followers."
	}
	
	// JSONにエンコードしてレスポンス
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}