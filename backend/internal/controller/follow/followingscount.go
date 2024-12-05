package followcontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)


func (fc *FollowController) GetFollowingsCount(w http.ResponseWriter, r *http.Request) {
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

	// Usecase層を呼び出してフォロワー数を更新・取得
	followingsCount, err := fc.followUsecase.GetFollowingsCount(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 成功レスポンスを返す
	response := map[string]int32{"followings_count": followingsCount}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}