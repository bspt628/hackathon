package likecontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (lc LikeController)GetLikeStatus(w http.ResponseWriter, r *http.Request) {
	firebaseUID := r.Header.Get("UserID")

	userID, err := lc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// HTTPリクエストのjsonをdecodeしてパラメータ「likeID」を取得
	var request struct {
		PostID string `json:"post_id"`
	}

	// リクエストのJSONデータを構造体にバインド
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエストの解析に失敗しました: %v", err), http.StatusBadRequest)
		return
	}

	// 必須パラメータをチェック
	if request.PostID == "" {
		http.Error(w, "PostID is required", http.StatusBadRequest)
		return 
	}

	liked, err := lc.likeUsecase.GetLikeStatus(context.Background(), userID, request.PostID)
	if err != nil {
		http.Error(w, fmt.Sprintf("いいねの存在確認に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(liked); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}