package likecontroller

import (
	"context"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func (lc *LikeController) CreateLike(w http.ResponseWriter, r *http.Request) {
	// FirebaseAuthMiddleware で設定された UserID を取得
	firebaseUID := r.Header.Get("UserID")
	if firebaseUID == "" {
		http.Error(w, "UserID missing in request context", http.StatusUnauthorized)
		return
	}

	userID, err := lc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("いいねするユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	postID := vars["id"]

	// 必須パラメータをチェック
	if postID == "" {
		http.Error(w, "PostID is required", http.StatusBadRequest)
		return 
	}



	// いいねを実行
	err = lc.likeUsecase.CreateLike(context.Background(), postID, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to like: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Liked successfully"}`))
}

