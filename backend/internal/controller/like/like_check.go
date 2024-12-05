package likecontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func (lc LikeController)GetLikeStatus(w http.ResponseWriter, r *http.Request) {
	firebaseUID := r.Header.Get("UserID")

	userID, err := lc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	postID := vars["id"]

	liked, err := lc.likeUsecase.GetLikeStatus(context.Background(), userID, postID)
	if err != nil {
		http.Error(w, fmt.Sprintf("いいねの存在確認に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]bool{"like_status": liked}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}