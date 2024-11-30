
package followcontroller

import (
	"context"
	"fmt"
	"net/http"
	"encoding/json"
)


func (fc *FollowController) GetFollowings(w http.ResponseWriter, r *http.Request) {
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
	followings, err := fc.followUsecase.GetFollowings(context.Background(), ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get followings: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"followings": followings,
	}
	
	// フォロワーがいない場合にメッセージを追加
	if len(followings) == 0 {
		response["followings"] = []interface{}{}
		response["message"] = "This user has no followings."
	}
	
	// JSONにエンコードしてレスポンス
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}