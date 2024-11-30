package followcontroller

import (
	"context"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func (fc *FollowController) RemoveFollow(w http.ResponseWriter, r *http.Request) {
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
	followingid := vars["id"]

	// 必須パラメータをチェック
	if followingid == "" {
		http.Error(w, "followingid is required", http.StatusBadRequest)
		return 
	}


	// フォローを実行
	err = fc.followUsecase.RemoveFollow(context.Background(), followingid, followerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to remove follow: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Follow removed successfully"}`))


}