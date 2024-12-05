package followcontroller

import (
	"context"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func (fc *FollowController) GetFollowStatus(w http.ResponseWriter, r *http.Request) {
	// FirebaseAuthMiddleware で設定された UserID を取得
	firebaseUID := r.Header.Get("UserID")
	if firebaseUID == "" {
		http.Error(w, "UserID missing in request context", http.StatusUnauthorized)
		return
	}

	followerID, err := fc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("フォローするユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}
	// URLパスからパラメータ「username」を取得
	vars := mux.Vars(r)
	followingID := vars["id"]

	// 必須パラメータをチェック
	if followingID == "" {
		http.Error(w, "followingid is required", http.StatusBadRequest)
		return 
	}


	// フォローを実行
	result, err := fc.followUsecase.GetFollowStatus(context.Background(), followingID, followerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get follow status: %v", err), http.StatusInternalServerError)
		return
	}

	// フォロー状態を返す
	response := map[string]bool{"follow_status": result}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}


}