package likecontroller

import (
	"context"
	"fmt"
	"net/http"
	"encoding/json"
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



	// いいねを実行
	err = lc.likeUsecase.CreateLike(context.Background(), request.PostID, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to like: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Liked successfully"}`))
}

